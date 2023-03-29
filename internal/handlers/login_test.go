package handlers_test

import (
	"bytes"
	"encoding/json"
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func get_body(password string) *bytes.Buffer {
	payload, _ := json.Marshal(map[string]string{
		"password": password,
	})
	body := bytes.NewBuffer(payload)
	return body
}

func TestLoginRoute(t *testing.T) {
	api := test.SetupApi()
	route := "/api/login"

	t.Run("login endpoint returns token", func(t *testing.T) {
		w := httptest.NewRecorder()

		password := configs.SetupConfigs().PASSWORD
		body := get_body(password)

		req, _ := http.NewRequest("POST", route, body)
		api.ServeHTTP(w, req)

		assert.Equal(t, w.Code, 200)

		response := make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Condition(t, func() bool {
			return len(response["access_token"]) > 0
		})
	})

	t.Run("login endpoint wrong password", func(t *testing.T) {
		w := httptest.NewRecorder()

		password := "wrong_password"
		body := get_body(password)

		req, _ := http.NewRequest("POST", route, body)
		api.ServeHTTP(w, req)

		assert.Equal(t, w.Code, 401)

		response := make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, response["error"], "unauthorized")
	})
}
