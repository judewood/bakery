package products

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/judewood/bakery/consts"
	"github.com/judewood/bakery/myfmt"
)

func TestGetProducts(t *testing.T) {
	mockProductService := NewMockProductService()
	mockProductService.On("GetAvailableProducts").Return(products)
	url := consts.Local_Base_Url + "/products"

	got := []Product{}
	err := getJson(url, &got)
	if err != nil {
		myfmt.Fatalf(t, "error while decoding available products to json format: %v", err)
	} 
	t.Logf("just got %v", got)
	if !reflect.DeepEqual(got, products) {
		myfmt.Fatalf(t, "\nWanted: %v. \n Got %v", products, got)
	}
}

func getJson(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
