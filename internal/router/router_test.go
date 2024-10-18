package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/judewood/bakery/utils/testutils"
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

			t.Log("Then I get 200 status code returned")
			{
				if w.Code == http.StatusOK {
					testutils.Passed(t)
				} else {
					testutils.Failed(t, w.Code)
				}
				t.Log("And I get response of 'pong'")
				{
					respBytes, err := io.ReadAll(w.Body)
					if err != nil {
						testutils.FailedToReadResponse(t, err)
					}
					if string(respBytes) != "pong" {
						testutils.Failed(t, string(respBytes))
					}
				}
			}
		}
	}
}
