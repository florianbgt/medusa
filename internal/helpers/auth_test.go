package helpers_test

import (
	"errors"
	"florianbgt/medusa/internal/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	t.Run("generates token pair", func(t *testing.T) {
		type testCase struct {
			key      string
			is_valid bool
		}

		for _, scenario := range []testCase{
			{
				key:      "valid_key",
				is_valid: true,
			},
			{
				key:      "invalid_key",
				is_valid: false,
			},
		} {
			token_pair, error := helpers.GenerateTokenPair(scenario.key)

			assert.Equal(t, error, nil)
			assert.Condition(t, func() bool {
				return len(token_pair.Access) > 0
			})
			assert.Condition(t, func() bool {
				return len(token_pair.Refresh) > 0
			})

			assert.Equal(t, helpers.IsTokenValid(token_pair.Access, "valid_key"), scenario.is_valid)
			assert.Equal(t, helpers.IsTokenValid(token_pair.Refresh, "valid_key"), scenario.is_valid)
		}
	})

	t.Run("get token from header", func(t *testing.T) {
		type testCase struct {
			header string
			token  string
			err    error
		}

		for _, scenario := range []testCase{
			{
				header: "",
				token:  "",
				err:    errors.New("bad header value"),
			},
			{
				header: "Bearer",
				token:  "",
				err:    errors.New("bad header format"),
			},
			{
				header: "Bearer token",
				token:  "token",
				err:    nil,
			},
		} {
			header := scenario.header
			token, err := helpers.GetTokenFromHeader(header)

			assert.Equal(t, err, scenario.err)
			assert.Equal(t, token, scenario.token)
		}
	})
}
