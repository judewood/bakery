package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/judewood/bakery/myfmt"
)

func TestPingRoute(t *testing.T) {
	t.Log("Given that I have the http server running")
	{
		router := SetupRouter()
		w := httptest.NewRecorder()

		t.Log("When I hit the ping endpoint")
		{
			req, _ := http.NewRequest("GET", "/ping", nil)
			router.ServeHTTP(w, req)

			t.Log("Then I get response of 'pong'")
			{
				if w.Code != http.StatusOK {
					t.Errorf("%sUnexpected http status.\nWant:\t%v\nGot:\t%v", myfmt.ThumbsDown,http.StatusOK, w.Code)
				}
				respBytes, err := io.ReadAll(w.Body)
				if err != nil {
					t.Errorf( "%sFailed to read response body bytes. %v",myfmt.ThumbsDown, err)
				}
				if string(respBytes) != "pong" {
					t.Errorf( "%sUnexpected response body.\nWant:\t%s\nGot:\t%s", myfmt.ThumbsDown, "pong", string(respBytes))
				}
			}
		}
	}
}
