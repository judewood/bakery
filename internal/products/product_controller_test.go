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
	t.Log("Given that I have the http server running")
	{
		t.Log("When there are  available products")
		{
			setUpGetProductsTest()
			mockProductService.On("GetAvailableProducts").Return(sampleProducts, nil)
			
			rtr.ServeHTTP(recorder, req)

			t.Log("Then I get 200 status")
			{
				if recorder.Code == http.StatusOK {
					t.Log(myfmt.ThumbsUp)
				} else {
					t.Errorf("\n%s Got  %v", myfmt.ThumbsDown, recorder.Code)
				}
			}
			t.Log("And I get the sample products")
			{
				responseData, _ := io.ReadAll(recorder.Body)
				gotProduct := []Product{}
				json.Unmarshal(responseData, &gotProduct)
				if reflect.DeepEqual(gotProduct, sampleProducts) {
					t.Log(myfmt.ThumbsUp)
				} else {
					t.Errorf("\n%s Got: %v", myfmt.ThumbsDown, recorder.Body)
				}
			}
		}

		t.Log("When there are no available products")
		{
			setUpGetProductsTest()
			mockProductService.On("GetAvailableProducts").Return([]Product{}, nil)
			
			rtr.ServeHTTP(recorder, req)

			t.Log("Then I get status 204")
			{
				if recorder.Code == http.StatusNoContent {
					t.Log(myfmt.ThumbsUp)
				} else {
					t.Errorf("\n%s Got %v", myfmt.ThumbsDown, recorder.Code)
				}
			}
			t.Log("And I get no available products")
			{
				responseData, _ := io.ReadAll(recorder.Body)
				gotProduct := []Product{}
				json.Unmarshal(responseData, &gotProduct)
				if reflect.DeepEqual(gotProduct, []Product{}) {
					t.Log(myfmt.ThumbsUp)
				} else {
					t.Errorf("\n%s Got: %v", myfmt.ThumbsDown, recorder.Body)
				}
			}
		}
	}
}

func TestAddProduct(t *testing.T) {
	t.Log("Given that I have the http server running")
	{
		t.Log("When I post a valid product")
		{
			setUpAddProductsTest()
			jsonBody, _ := json.Marshal(&sampleProducts[0])
			bodyReader := bytes.NewReader(jsonBody)
			req, _ := http.NewRequest("POST", "/products", bodyReader)
			mockProductService.On("Add").Return(sampleProducts[0], nil)

			rtr = router.AddProduct(rtr, controller.Add)
			rtr.ServeHTTP(recorder, req)
			
			t.Log("Then I get 204 status")
			{
				if recorder.Code == http.StatusCreated {
					t.Log(myfmt.ThumbsUp)

				} else {
					t.Errorf("\n%s Got %v", myfmt.ThumbsDown, recorder.Code)
				}
			}
			t.Log("And sampleProduct[0] is returned")
			{
				responseData, _ := io.ReadAll(recorder.Body)
				gotProduct := Product{}
				json.Unmarshal(responseData, &gotProduct)
				if reflect.DeepEqual(gotProduct, sampleProducts[0]) {
					t.Log(myfmt.ThumbsUp)
				} else {
					t.Errorf("\n%s Got: %v", myfmt.ThumbsDown, recorder.Body)
				}
			}
		}
		t.Log("When I post an invalid product")
		{
			setUpAddProductsTest()
			invalid := Product{ Name: "", RecipeID: "1"}
			jsonBody, _ := json.Marshal(&invalid)
			bodyReader := bytes.NewReader(jsonBody)
			req, _ := http.NewRequest("POST", "/products", bodyReader)
			mockProductService.On("Add").Return(Product{}, errors.New(MissingRequired))
			rtr = router.AddProduct(rtr, controller.Add)

			rtr.ServeHTTP(recorder, req)

			t.Log("Then I get 400 status")
			{
				if recorder.Code == http.StatusBadRequest {
					t.Log(myfmt.ThumbsUp)

				} else {
					t.Errorf("\n%s Got %v", myfmt.ThumbsDown, recorder.Code)
				}
			}
			t.Log("And nothing is returned")
			{
				responseData, _ := io.ReadAll(recorder.Body)
				gotProduct := Product{}
				json.Unmarshal(responseData, &gotProduct)
				if reflect.DeepEqual(gotProduct, Product{}) {
					t.Log(myfmt.ThumbsUp)
				} else {
					t.Errorf("\n%s Got: %v", myfmt.ThumbsDown, gotProduct)
				}
			}
		}
	}
}
