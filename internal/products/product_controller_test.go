package products

import (
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
	r = router.AddRouteGetProducts(r, controller.GetProducts)
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
