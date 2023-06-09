package password_change_test

import (
	"bytes"
	"encoding/json"
	"florianbgt/medusa/internal/helpers"
	"florianbgt/medusa/internal/models/password_model"
	"florianbgt/medusa/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordChangeRoute(t *testing.T) {
	api := test.SetupApi()
	route := "/api/password/change"
	db := test.Setupdb()

	var Password password_model.Password

	configs := test.SetupConfigs()

	Password.Setup(db, configs.DEFAULT_PASSWORD, configs.API_KEY)

	t.Run("change password", func(t *testing.T) {
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
				err:     "Key: 'OldPassword' Error:Field validation for 'OldPassword' failed on the 'required' tag\nKey: 'Password' Error:Field validation for 'Password' failed on the 'required' tag\nKey: 'Password2' Error:Field validation for 'Password2' failed on the 'required' tag",
			},
			{
				payload: map[string]string{
					"old_password": configs.DEFAULT_PASSWORD,
					"password":     "Newpassword/123",
					"password2":    "Newpassword/123",
				},
				status:  http.StatusOK,
				success: true,
				err:     "",
			},
			{
				payload: map[string]string{
					"old_password": "wrongpassword",
					"password":     "Newpassword/123",
					"password2":    "Newpassword/123",
				},
				status:  http.StatusBadRequest,
				success: false,
				err:     "old_password_incorrect",
			},
			{
				payload: map[string]string{
					"old_password": configs.DEFAULT_PASSWORD,
					"password":     "Newpassword/123",
					"password2":    "password2_does_not_match",
				},
				status:  http.StatusBadRequest,
				success: false,
				err:     "password2_does_not_match",
			},
			{
				payload: map[string]string{
					"old_password": configs.DEFAULT_PASSWORD,
					"password":     "newpassword/123",
					"password2":    "newpassword/123",
				},
				status:  http.StatusBadRequest,
				success: false,
				err:     "invalid_password",
			},
		} {
			// reset password
			Password.UpdatePassword(db, "Password/123", configs.API_KEY)

			w := httptest.NewRecorder()

			payload, _ := json.Marshal(scenario.payload)
			body := bytes.NewBuffer(payload)
			token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

			req, _ := http.NewRequest("POST", route, body)
			req.Header = http.Header{
				"Authorization": []string{"Bearer " + token_pair.Access},
			}
			api.ServeHTTP(w, req)

			assert.Equal(t, scenario.status, w.Code)

			response := make(map[string]string)
			json.Unmarshal(w.Body.Bytes(), &response)

			if !scenario.success {
				assert.Equal(t, scenario.err, response["error"])
			}
		}
	})
}
