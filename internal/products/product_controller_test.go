package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func setUpGetProductsTest() {
	setUp()
	req, _ = http.NewRequest("GET", "/products", nil)
	rtr = router.GetProducts(rtr, controller.GetProducts)
}

func setUpAddProductsTest() {
	setUp()
}

func TestGetProducts(t *testing.T) {
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
					setUpGetProductsTest()
					mockProductService.On("GetAvailableProducts").Return(test.body, nil)

					rtr.ServeHTTP(recorder, req)

					t.Logf("Then I get %d status", test.status)
					{
						if recorder.Code == test.status {
							t.Log(myfmt.ThumbsUp)
						} else {
							t.Errorf("\n%s Got  %v", myfmt.ThumbsDown, recorder.Code)
						}
					}
					t.Logf("And I get %v", test.body)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						gotProduct := []Product{}
						json.Unmarshal(responseData, &gotProduct)
						if reflect.DeepEqual(gotProduct, test.body) {
							t.Log(myfmt.ThumbsUp)
						} else {
							t.Errorf("\n%s Got: %v", myfmt.ThumbsDown, recorder.Body)
						}
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

func TestAddProduct(t *testing.T) {
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
					setUpAddProductsTest()
					jsonBody, _ := json.Marshal(&test.reqProduct)
					bodyReader := bytes.NewReader(jsonBody)
					req, _ := http.NewRequest("POST", "/products", bodyReader)
					mockProductService.On("Add").Return(test.respProduct, test.err)

					rtr = router.AddProduct(rtr, controller.Add)
					rtr.ServeHTTP(recorder, req)

					t.Logf("Then I get %d status", test.status)
					{
						if recorder.Code == test.status {
							t.Log(myfmt.ThumbsUp)

						} else {
							t.Errorf("\n%s Got %v", myfmt.ThumbsDown, recorder.Code)
						}
					}
					t.Logf("And %v is returned", test.respProduct)
					{
						responseData, _ := io.ReadAll(recorder.Body)
						if (test.respProduct == Product{} && string(responseData) == `""`) {
							t.Log(myfmt.ThumbsUp)
						} else {
							gotProduct := Product{}
							err := json.Unmarshal(responseData, &gotProduct)
							if err != nil {
								t.Errorf("\n%s failed to deserialise the response: %v", myfmt.ThumbsDown, recorder.Body)
							}
							if reflect.DeepEqual(gotProduct, test.reqProduct) {
								t.Log(myfmt.ThumbsUp)
							} else {
								t.Errorf("\n%s Got: %v", myfmt.ThumbsDown, recorder.Body)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}

	}
}
