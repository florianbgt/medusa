package refresh_test

import (
	"bytes"
	"encoding/json"
	"florianbgt/medusa/internal/helpers"
	"florianbgt/medusa/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefreshTokenRoute(t *testing.T) {
	api := test.SetupApi()
	route := "/api/token/refresh"

	configs := test.SetupConfigs()

	token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

	t.Run("refresh token", func(t *testing.T) {
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
				err:     "Key: 'Refresh' Error:Field validation for 'Refresh' failed on the 'required' tag",
			},
			{
				payload: map[string]string{
					"refresh": "invalid_token",
				},
				status:  http.StatusUnauthorized,
				success: false,
				err:     "unauthorized",
			},
			{
				payload: map[string]string{
					"refresh": token_pair.Refresh,
				},
				status:  http.StatusOK,
				success: true,
				err:     "",
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
				assert.Equal(t, response["error"], scenario.err)
			}
		}
	})
}
