package products

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/judewood/bakery/internal/router"
	"github.com/judewood/bakery/myfmt"
)

func TestGetProducts(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	mockProductService := NewMockProductService()
	mockProductService.On("GetAvailableProducts").Return(sampleProducts, nil)

	controller := NewProductController(mockProductService)
	r = router.GetProducts(r, controller.GetProducts)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		myfmt.Errorf(t, "\nExpected status %v.\nGot  %v", http.StatusOK, w.Code)
	}
	responseData, _ := io.ReadAll(w.Body)
	gotProduct := []Product{}
	json.Unmarshal(responseData, &gotProduct)
	if !reflect.DeepEqual(gotProduct, sampleProducts) {
		myfmt.Errorf(t, "Failed to get products. \nWanted: %v\nGot: %v", sampleProducts, w.Body)
	}
}

func TestAddProduct(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder()

	//jsonBody := []byte(`{"client_message": "hello, server!"}`)
	jsonBody, _ := json.Marshal(&sampleProducts[0])

	bodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "/products", bodyReader)
	mockProductService := NewMockProductService()
	mockProductService.On("Add").Return(sampleProducts[0], nil)

	controller := NewProductController(mockProductService)
	r = router.AddProduct(r, controller.Add)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		myfmt.Errorf(t, "\nExpected status %v.\nGot  %v", http.StatusCreated, w.Code)
	}
	responseData, _ := io.ReadAll(w.Body)
	gotProduct := Product{}
	json.Unmarshal(responseData, &gotProduct)
	if !reflect.DeepEqual(gotProduct, sampleProducts[0]) {
		myfmt.Errorf(t, "Failed to add product. \nWanted: %v\nGot: %v", sampleProducts[0], w.Body)
	}
}
