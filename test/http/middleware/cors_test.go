package middleware

import (
	"net/http"
	"testing"

	test "github.com/bayusamudra5502/go-backend-template/test/http"
	"github.com/stretchr/testify/assert"
)

func TestPreflight(t *testing.T) {
	t.Run("PreflightShouldBeOk", func(t *testing.T) {
		res, _, err := test.ExecuteJSON(test.RequestData{
			Method: "OPTIONS",
			Endpoint: "/ping",
			Headers: map[string]string{
				"Access-Control-Request-Method": "GET",
				"Access-Control-Request-Headers": "accept, origin, authorization, content-type, referer",
				"Origin": "https://inkubatorit.com",
			},
		})

		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, res.Header.Get("Access-Control-Allow-Origin"), "https://inkubatorit.com")
		assert.Contains(t, res.Header.Get("Access-Control-Allow-Methods"), "GET")
	})

	t.Run("PreflightShouldBeOkOnNotFound", func(t *testing.T) {
		res, _, err := test.ExecuteJSON(test.RequestData{
			Method: "OPTIONS",
			Endpoint: "/not-found",
			Headers: map[string]string{
				"Access-Control-Request-Method": "GET",
				"Access-Control-Request-Headers": "accept, origin, authorization, content-type, referer",
				"Origin": "https://inkubatorit.com",
			},
		})

		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, res.Header.Get("Access-Control-Allow-Origin"), "https://inkubatorit.com")
		assert.Contains(t, res.Header.Get("Access-Control-Allow-Methods"), "GET")
	})

	t.Run("PreflightNotAllowedPatchMethod", func(t *testing.T) {
		res, _, err := test.ExecuteJSON(test.RequestData{
			Method: "OPTIONS",
			Endpoint: "/ping",
			Headers: map[string]string{
				"Access-Control-Request-Method": "PATCH",
				"Access-Control-Request-Headers": "accept, origin, authorization, content-type, referer",
				"Origin": "https://inkubatorit.com",
			},
		})

		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, res.Header.Get("Access-Control-Allow-Origin"), "")
		assert.Equal(t, res.Header.Get("Access-Control-Allow-Methods"), "")
	})
}