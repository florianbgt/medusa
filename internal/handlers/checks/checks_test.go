package checks_test

import (
	"florianbgt/medusa/internal/helpers"
	"florianbgt/medusa/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthyRoute(t *testing.T) {
	api := test.SetupApi()
	route := "/api/healthy"

	t.Run("health endpoint returns healthy", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", route, nil)
		api.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.String(), "\"healthy\"")
	})
}

func TestAuthenticatedRoute(t *testing.T) {
	api := test.SetupApi()
	route := "/api/authenticated"

	configs := test.SetupConfigs()

	t.Run("non authenticated returns 401", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", route, nil)
		api.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusUnauthorized)
	})

	t.Run("authenticated returns 200", func(t *testing.T) {
		w := httptest.NewRecorder()

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

		req, _ := http.NewRequest("GET", route, nil)
		req.Header = http.Header{
			"Authorization": []string{"Bearer " + token_pair.Access},
		}

		api.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
	})
}
