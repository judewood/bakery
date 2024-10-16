package products

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/judewood/bakery/errorutils"
	"github.com/judewood/bakery/internal/router"
	"github.com/judewood/bakery/testutils"
)

var recorder *httptest.ResponseRecorder
var rtr *gin.Engine
var mockProductService *MockProductService
var controller *ProductController
var req *http.Request

func setUp() {
	rtr = router.SetupRouter()
	recorder = httptest.NewRecorder()
	mockProductService = NewMockProductService()
	controller = NewProductController(mockProductService)
}

func TestProductControllerGetAll(t *testing.T) {
	testCases := []struct {
		name   string
		body   []Product
		status int
	}{
		{name: "some", body: sampleProducts, status: http.StatusOK},
		{name: "zero", body: []Product{}, status: http.StatusNoContent},
	}

	t.Log("Given that I have the http server running")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("\n test %d: When I have %s products", i, test.name)
				{
					setUp()
					req, _ = http.NewRequest("GET", "/products", nil)
					rtr = router.GetProducts(rtr, controller.GetProducts)
					mockProductService.On("GetAll").Return(test.body, nil)

					rtr.ServeHTTP(recorder, req)

					t.Logf("Then I get %d status", test.status)
					{
						if recorder.Code == test.status {
							testutils.Passed(t)
						} else {
							testutils.Failed(t, recorder.Code)
						}
					}
					t.Logf("And I get %v", test.body)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						gotProduct := []Product{}
						json.Unmarshal(responseData, &gotProduct)
						if reflect.DeepEqual(gotProduct, test.body) {
							testutils.Passed(t)
						} else {
							testutils.Failed(t, recorder.Body)
						}
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

func TestProductControllerGet(t *testing.T) {
	type testCase struct {
		name        string
		reqId       string
		respProduct Product
		status      int
		err         error
	}
	testCases := []testCase{
		{name: "product exists", reqId: strings.ToLower(sampleProducts[2].Name), respProduct: sampleProducts[2], status: http.StatusOK, err: nil},
		{name: "product does not exist", reqId: "missing", respProduct: Product{}, status: http.StatusNoContent, err: errorutils.NotFoundError("missing")},
		{name: "requested id is invalid", reqId: "a", respProduct: Product{}, status: http.StatusBadRequest, err: errorutils.ErrorMissingID},
	}
	t.Log("Given that I have the http server running")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("Test: %d When the %s", i, test.name)
				{
					setUp()
					req, _ := http.NewRequest("GET", "/product/"+test.reqId, nil)

					mockProductService.On("Get").Return(test.respProduct, test.err)

					rtr = router.GetProduct(rtr, controller.Get)
					rtr.ServeHTTP(recorder, req)

					t.Logf("Then I get %d status", test.status)
					{
						if recorder.Code == test.status {
							testutils.Passed(t)

						} else {
							testutils.Failed(t, recorder.Code)
						}
					}
					t.Logf("And %v is returned", test.respProduct)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						if test.err != nil && (string(responseData) == `""` || string(responseData) == "") {
							testutils.Passed(t)
						} else {
							gotProduct := Product{}
							err := json.Unmarshal(responseData, &gotProduct)
							if err != nil {
								testutils.Failed(t, recorder.Body)
							}
							if reflect.DeepEqual(gotProduct, test.respProduct) {
								testutils.Passed(t)
							} else {
								testutils.Failed(t, gotProduct)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}

	}
}

func TestProductControllerAdd(t *testing.T) {
	type testCase struct {
		name     string
		reqBody  any
		respBody Product
		status   int
		err      error
	}
	testCases := []testCase{
		{name: "valid", reqBody: sampleProducts[0], respBody: sampleProducts[0], status: http.StatusCreated, err: nil},
		{name: "empty product", reqBody: Product{}, respBody: Product{}, status: http.StatusBadRequest, err: errorutils.ErrorMissingRequired},
		{name: "missing product name", reqBody: Product{Name: "", RecipeID: "4"}, respBody: Product{}, status: http.StatusBadRequest, err: errorutils.ErrorMissingRequired},
		{name: "missing recipeID", reqBody: Product{Name: "product name", RecipeID: ""}, respBody: Product{}, status: http.StatusBadRequest, err: errorutils.ErrorMissingRequired},
		{name: "not a product", reqBody: "Not a product", respBody: Product{}, status: http.StatusBadRequest, err: errorutils.ErrorMissingRequired},
	}
	t.Log("Given that I have the http server running")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("Test: %d When I post a product that is %s", i, test.name)
				{
					setUp()
					jsonBody, _ := json.Marshal(&test.reqBody)
					bodyReader := bytes.NewReader(jsonBody)
					req, _ := http.NewRequest("POST", "/product", bodyReader)
					mockProductService.On("Add").Return(test.respBody, test.err)

					rtr = router.AddProduct(rtr, controller.Add)
					rtr.ServeHTTP(recorder, req)

					t.Logf("Then I get %d status", test.status)
					{
						if recorder.Code == test.status {
							testutils.Passed(t)

						} else {
							testutils.Failed(t, recorder.Code)
						}
					}
					t.Logf("And %v is returned", test.respBody)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						if (test.respBody == Product{} && string(responseData) == `""`) {
							testutils.Passed(t)
						} else {
							gotProduct := Product{}
							err := json.Unmarshal(responseData, &gotProduct)
							if err != nil {
								testutils.FailedToDeserialise(t, recorder.Body)
							}
							if reflect.DeepEqual(gotProduct, test.reqBody) {
								testutils.Passed(t)
							} else {
								testutils.Failed(t, recorder.Body)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

func TestProductControllerUpdate(t *testing.T) {
	type testCase struct {
		name     string
		reqBody  any
		respBody Product
		status   int
		err      error
	}
	testCases := []testCase{
		{name: "valid", reqBody: sampleProducts[0], respBody: sampleProducts[0], status: http.StatusOK, err: nil},
		{name: "empty product", reqBody: Product{}, respBody: Product{}, status: http.StatusBadRequest, err: errorutils.ErrorMissingRequired},
		{name: "missing product name", reqBody: Product{Name: "", RecipeID: "4"}, respBody: Product{}, status: http.StatusBadRequest, err: errorutils.ErrorMissingRequired},
		{name: "missing recipeID", reqBody: Product{Name: "product name", RecipeID: ""}, respBody: Product{}, status: http.StatusBadRequest, err: errorutils.ErrorMissingRequired},
		{name: "not a product", reqBody: "Not a product", respBody: Product{}, status: http.StatusBadRequest, err: errorutils.ErrorMissingRequired},
		{name: "not found", reqBody: Product{Name: "missing", RecipeID: "4"}, respBody: Product{}, status: http.StatusNoContent, err: errorutils.NotFoundError("missing")},
	}
	t.Log("Given that I have the http server running")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("Test: %d When I update a product that is %s", i, test.name)
				{
					setUp()
					jsonBody, _ := json.Marshal(&test.reqBody)
					bodyReader := bytes.NewReader(jsonBody)
					req, _ := http.NewRequest("PUT", "/product", bodyReader)
					mockProductService.On("Update").Return(test.respBody, test.err)

					rtr = router.UpdateProduct(rtr, controller.Update)
					rtr.ServeHTTP(recorder, req)

					t.Logf("Then I get %d status", test.status)
					{
						if recorder.Code == test.status {
							testutils.Passed(t)

						} else {
							testutils.Failed(t, recorder.Code)
						}
					}
					t.Logf("And %v is returned", test.respBody)
					{
						responseData, _ := io.ReadAll(recorder.Body)

						if test.err != nil && (string(responseData) == `""` || string(responseData) == "") {
							testutils.Passed(t)
						} else {
							gotProduct := Product{}
							err := json.Unmarshal(responseData, &gotProduct)
							if err != nil {
								testutils.FailedToDeserialise(t, recorder.Body)
							}
							if reflect.DeepEqual(gotProduct, test.reqBody) {
								testutils.Passed(t)
							} else {
								testutils.Failed(t, recorder.Body)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}

	}
}

func TestProductControllerDelete(t *testing.T) {
	type testCase struct {
		name        string
		reqId       string
		respProduct Product
		status      int
		err         error
	}
	testCases := []testCase{

		{name: "product exists", reqId: strings.ToLower(sampleProducts[2].Name), respProduct: sampleProducts[2], status: http.StatusOK, err: nil},
		{name: "product does not exist", reqId: "missing", respProduct: Product{}, status: http.StatusNoContent, err: errorutils.NotFoundError("missing")},
		{name: "requested id is invalid", reqId: "a", respProduct: Product{}, status: http.StatusBadRequest, err: errorutils.ErrorMissingID},
	}
	t.Log("Given that I want to delete a product")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("Test: %d When the %s", i, test.name)
				{
					setUp()
					req, _ := http.NewRequest("DELETE", "/product/"+test.reqId, nil)

					mockProductService.On("Delete").Return(test.respProduct, test.err)

					rtr = router.DeleteProduct(rtr, controller.Delete)
					rtr.ServeHTTP(recorder, req)

					t.Logf("Then I get %d status", test.status)
					{
						if recorder.Code == test.status {
							testutils.Passed(t)

						} else {
							testutils.Failed(t, recorder.Code)
						}
					}
					t.Logf("And %v is returned", test.respProduct)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						if (test.respProduct == Product{} && string(responseData) == `""` || len(responseData) == 0) {
							testutils.Passed(t)
						} else {
							got := Product{}
							err := json.Unmarshal(responseData, &got)
							if err != nil {
								testutils.FailedToDeserialise(t, recorder.Body)
							}
							if reflect.DeepEqual(got, test.respProduct) {
								testutils.Passed(t)
							} else {
								testutils.Failed(t, got)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}

	}
}
