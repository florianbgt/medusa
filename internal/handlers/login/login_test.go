package login_test

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

func TestLoginRoute(t *testing.T) {
	api := test.SetupApi()
	route := "/api/login"

	t.Run("login", func(t *testing.T) {
		type testCase struct {
			payload map[string]string
			status  int
			success bool
			err     string
		}

		for _, scenario := range []testCase{
			{
				payload: map[string]string{},
				status:  http.StatusBadRequest,
				success: false,
				err:     "Key: 'Password' Error:Field validation for 'Password' failed on the 'required' tag",
			},
			{
				payload: map[string]string{
					"password": configs.SetupConfigs().DEFAULT_PASSWORD,
				},
				status:  http.StatusOK,
				success: true,
				err:     "",
			},
			{
				payload: map[string]string{
					"password": "wrong_password",
				},
				status:  http.StatusBadRequest,
				success: false,
				err:     "password_incorrect",
			},
		} {
			w := httptest.NewRecorder()

			payload, _ := json.Marshal(scenario.payload)
			body := bytes.NewBuffer(payload)

			req, _ := http.NewRequest("POST", route, body)
			api.ServeHTTP(w, req)

			assert.Equal(t, scenario.status, w.Code)

			response := make(map[string]string)
			json.Unmarshal(w.Body.Bytes(), &response)

			if scenario.success {
				assert.Condition(t, func() bool {
					return len(response["access_token"]) > 0
				})
				assert.Condition(t, func() bool {
					return len(response["refresh_token"]) > 0
				})
			} else {
				assert.Equal(t, scenario.err, response["error"])
			}
		}
	})
}
