package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/judewood/bakery/internal/router"
	"github.com/judewood/bakery/myfmt"
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
							passed(t)
						} else {
							failed(t,recorder.Code)
						}
					}
					t.Logf("And I get %v", test.body)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						gotProduct := []Product{}
						json.Unmarshal(responseData, &gotProduct)
						if reflect.DeepEqual(gotProduct, test.body) {
							passed(t)
						} else {
							failed(t, recorder.Body)
						}
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

func TestProductControllerGet(t *testing.T) {
	invalidReqErr := errors.New(MissingID)
	notFoundError := fmt.Errorf(NotFound, "missing")
	type testCase struct {
		name        string
		reqId       string
		respProduct Product
		status      int
		err         error
	}
	testCases := []testCase{
		{name: "product exists", reqId: strings.ToLower(sampleProducts[2].Name), respProduct: sampleProducts[2], status: http.StatusOK, err: nil},
		{name: "product does not exist", reqId: "missing", respProduct: Product{}, status: http.StatusNoContent, err: notFoundError},
		{name: "requested id is invalid", reqId: "a", respProduct: Product{}, status: http.StatusBadRequest, err: invalidReqErr},
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
							passed(t)

						} else {
							failed(t, recorder.Code)
						}
					}
					t.Logf("And %v is returned", test.respProduct)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						if (test.respProduct == Product{} && string(responseData) == `""` || len(responseData) == 0) {
							passed(t)
						} else {
							gotProduct := Product{}
							err := json.Unmarshal(responseData, &gotProduct)
							if err != nil {
								failed(t ,recorder.Body)
							}
							if reflect.DeepEqual(gotProduct, test.respProduct) {
								passed(t)
							} else {
								failed(t, gotProduct)
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
	invalidReqErr := errors.New(MissingRequired)
	type testCase struct {
		name        string
		reqProduct  Product
		respProduct Product
		status      int
		err         error
	}
	testCases := []testCase{
		{name: "valid", reqProduct: sampleProducts[0], respProduct: sampleProducts[0], status: http.StatusCreated, err: nil},
		{name: "empty product", reqProduct: Product{}, respProduct: Product{}, status: http.StatusBadRequest, err: invalidReqErr},
		{name: "missing product name", reqProduct: Product{Name: "", RecipeID: "4"}, respProduct: Product{}, status: http.StatusBadRequest, err: invalidReqErr},
		{name: "missing recipeID", reqProduct: Product{Name: "product name", RecipeID: ""}, respProduct: Product{}, status: http.StatusBadRequest, err: invalidReqErr},
	}
	t.Log("Given that I have the http server running")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("Test: %d When I post a product that is %s", i, test.name)
				{
					setUp()
					jsonBody, _ := json.Marshal(&test.reqProduct)
					bodyReader := bytes.NewReader(jsonBody)
					req, _ := http.NewRequest("POST", "/product", bodyReader)
					mockProductService.On("Add").Return(test.respProduct, test.err)

					rtr = router.AddProduct(rtr, controller.Add)
					rtr.ServeHTTP(recorder, req)

					t.Logf("Then I get %d status", test.status)
					{
						if recorder.Code == test.status {
							passed(t)

						} else {
							failed(t, recorder.Code)
						}
					}
					t.Logf("And %v is returned", test.respProduct)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						if (test.respProduct == Product{} && string(responseData) == `""`) {
							passed(t)
						} else {
							gotProduct := Product{}
							err := json.Unmarshal(responseData, &gotProduct)
							if err != nil {
								failedToDeserialiseBody(t, recorder.Body)
							}
							if reflect.DeepEqual(gotProduct, test.reqProduct) {
								passed(t)
							} else {
								failed(t, recorder.Body)
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
	invalidReqErr := errors.New(MissingRequired)
	type testCase struct {
		name        string
		reqProduct  Product
		respProduct Product
		status      int
		err         error
	}
	testCases := []testCase{
		{name: "valid", reqProduct: sampleProducts[0], respProduct: sampleProducts[0], status: http.StatusOK, err: nil},
		{name: "empty product", reqProduct: Product{}, respProduct: Product{}, status: http.StatusBadRequest, err: invalidReqErr},
		{name: "missing product name", reqProduct: Product{Name: "", RecipeID: "4"}, respProduct: Product{}, status: http.StatusBadRequest, err: invalidReqErr},
		{name: "missing recipeID", reqProduct: Product{Name: "product name", RecipeID: ""}, respProduct: Product{}, status: http.StatusBadRequest, err: invalidReqErr},
	}
	t.Log("Given that I have the http server running")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("Test: %d When I update a product that is %s", i, test.name)
				{
					setUp()
					jsonBody, _ := json.Marshal(&test.reqProduct)
					bodyReader := bytes.NewReader(jsonBody)
					req, _ := http.NewRequest("PUT", "/product", bodyReader)
					mockProductService.On("Update").Return(test.respProduct, test.err)

					rtr = router.UpdateProduct(rtr, controller.Update)
					rtr.ServeHTTP(recorder, req)

					t.Logf("Then I get %d status", test.status)
					{
						if recorder.Code == test.status {
							passed(t)

						} else {
							failed(t, recorder.Code)
						}
					}
					t.Logf("And %v is returned", test.respProduct)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						if (test.respProduct == Product{} && string(responseData) == `""`) {
							passed(t)
						} else {
							gotProduct := Product{}
							err := json.Unmarshal(responseData, &gotProduct)
							if err != nil {
								failedToDeserialiseBody(t, recorder.Body)
							}
							if reflect.DeepEqual(gotProduct, test.reqProduct) {
								passed(t)
							} else {
								failed(t, recorder.Body)
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
	invalidReqErr := errors.New(MissingID)
	notFoundError := fmt.Errorf(NotFound, "missing")
	type testCase struct {
		name        string
		reqId       string
		respProduct Product
		status      int
		err         error
	}
	testCases := []testCase{

		{name: "product exists", reqId: strings.ToLower(sampleProducts[2].Name), respProduct: sampleProducts[2], status: http.StatusOK, err: nil},
		{name: "product does not exist", reqId: "missing", respProduct: Product{}, status: http.StatusNoContent, err: notFoundError},
		{name: "requested id is invalid", reqId: "a", respProduct: Product{}, status: http.StatusBadRequest, err: invalidReqErr},
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
							passed(t)

						} else {
							failed(t, recorder.Code)
						}
					}
					t.Logf("And %v is returned", test.respProduct)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						if (test.respProduct == Product{} && string(responseData) == `""` || len(responseData) == 0) {
							passed(t)
						} else {
							got := Product{}
							err := json.Unmarshal(responseData, &got)
							if err != nil {
								failedToDeserialiseBody(t, recorder.Body)
							}
							if reflect.DeepEqual(got, test.respProduct) {
								passed(t)
							} else {
								failed(t, got)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}

	}
}

func passed (t *testing.T) {
	t.Log(myfmt.ThumbsUp)
}

func failed(t *testing.T, got any) {
	t.Errorf("\n%s Got: %v", myfmt.ThumbsDown, got)
}

func failedToDeserialiseBody(t *testing.T, body any) {
	t.Errorf("\n%s failed to deserialise the response: %v", myfmt.ThumbsDown, body)
}
