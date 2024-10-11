package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/judewood/bakery/myfmt"
)

func TestPingRoute(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		myfmt.Fatalf(t, "Unexpected status when pinging server. \nWanted: %v \n Got: %v", http.StatusOK, w.Code)
	}
	if w.Body.String() != "pong" {
		myfmt.Fatalf(t, "Unexpected response body  when pinging server. \nWanted: %v \n Got: %v", "pong", w.Body.String())
	}
}
