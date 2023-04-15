package password_model_test

import (
	"errors"
	"florianbgt/medusa/internal/models/password_model"
	"florianbgt/medusa/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	db := test.Setupdb()

	configs := test.SetupConfigs()

	var Password password_model.Password

	t.Run("update password", func(t *testing.T) {
		Password.Setup(db, "Password/123", configs.API_KEY)

		type testCase struct {
			password string
			err      error
		}

		for _, scenario := range []testCase{
			{
				password: "Password/321",
				err:      nil,
			},
			{
				password: "",
				err:      errors.New("invalid_password"),
			},
			{
				password: "password/321",
				err:      errors.New("invalid_password"),
			},
			{
				password: "PASSWORD/321",
				err:      errors.New("invalid_password"),
			},
			{
				password: "Password/onetwothree",
				err:      errors.New("invalid_password"),
			},
			{
				password: "Password321",
				err:      errors.New("invalid_password"),
			},
		} {
			// reset password before each run
			Password.UpdatePassword(db, "Password/123", configs.API_KEY)

			err := Password.UpdatePassword(db, scenario.password, configs.API_KEY)

			assert.Equal(t, scenario.err, err)

			if err == nil {
				assert.True(t, password_model.CheckPasswordHash(scenario.password, configs.API_KEY, Password.Password))
			} else {
				assert.True(t, password_model.CheckPasswordHash("Password/123", configs.API_KEY, Password.Password))
			}
		}
	})
}
