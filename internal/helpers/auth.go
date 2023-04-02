package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type TokenPair struct {
	Access  string
	Refresh string
}

func GenerateTokenPair(api_key string) (TokenPair, error) {
	token_pair := TokenPair{}

	access_token := jwt.New(jwt.SigningMethodHS256)
	claims := access_token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	access_token_string, err := access_token.SignedString([]byte(api_key))
	if err != nil {
		return token_pair, errors.New("could not generate access token")
	}

	refresh_token := jwt.New(jwt.SigningMethodHS256)
	claims = refresh_token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	refresh_token_string, err := refresh_token.SignedString([]byte(api_key))
	if err != nil {
		return token_pair, errors.New("could not generate refresh token")
	}

	token_pair.Access = access_token_string
	token_pair.Refresh = refresh_token_string

	return token_pair, nil
}

func GetTokenFromHeader(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value")
	}

	bearer_token := strings.Split(header, " ")

	if len(bearer_token) != 2 {
		return "", errors.New("bad header format")
	}

	return bearer_token[1], nil
}

func IsTokenValid(token string, api_key string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, OK := token.Method.(*jwt.SigningMethodHMAC)
		if !OK {
			return nil, errors.New("bad signing method")
		}
		return []byte(api_key), nil
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func IsAuthCheck(c *gin.Context, api_key string) {
	auth_header := c.GetHeader("Authorization")

	token, err := GetTokenFromHeader(auth_header)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	if !IsTokenValid(token, api_key) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	c.Next()
}
