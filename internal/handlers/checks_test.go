package handlers_test

import (
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

		assert.Equal(t, w.Code, 200)
		assert.Equal(t, w.Body.String(), "healthy")
	})
}
