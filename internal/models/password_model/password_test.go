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

	var passwordInstance password_model.Password

	passwordInstance.Setup(db, "Password/123")
	t.Run("update password", func(t *testing.T) {
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
			passwordInstance.UpdatePassword(db, "Password/123")

			err := passwordInstance.UpdatePassword(db, scenario.password)

			excpectedPassword := "Password/123"
			if scenario.err == nil {
				excpectedPassword = scenario.password
			}

			assert.Equal(t, scenario.err, err)
			assert.Equal(t, excpectedPassword, passwordInstance.Password)
		}
	})

	t.Run("get password success", func(t *testing.T) {
		passwordInstance.UpdatePassword(db, "Password/123")

		password, err := passwordInstance.GetPassword(db)

		assert.Equal(t, nil, err)
		assert.Equal(t, "Password/123", password)
	})

	t.Run("get password fail", func(t *testing.T) {
		db.Delete(&passwordInstance)

		password, err := passwordInstance.GetPassword(db)

		assert.Equal(t, nil, err)
		assert.Equal(t, "Password/123", password)
	})
}
