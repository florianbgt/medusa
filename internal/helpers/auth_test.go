package helpers_test

import (
	"florianbgt/medusa/internal/helpers"
	testing "testing"

	assert "github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	t.Run("generates a valid token", func(t *testing.T) {
		token, error := helpers.GenerateToken("test_key", 1)

		assert.Equal(t, error, nil)
		assert.Condition(t, func() bool {
			return len(token) > 0
		})
	})
}

// func TestIsAuthCheck(t *testing.T) {
// 	t.Run("valid in request process", func(t *testing.T) {
// 		api_key := "test_key"
// 		token, _ := helpers.GenerateToken(api_key, 1)

// 	})
// }
