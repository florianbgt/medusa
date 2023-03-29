package helpers

import (
	errors "errors"
	http "net/http"
	strings "strings"
	time "time"

	gin "github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(api_key string, expiration int) (string, error) {
	expirationTime := time.Now().Add(
		time.Duration(expiration) * time.Second,
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expiration": jwt.NewNumericDate(expirationTime),
	})

	jwtKey := []byte(api_key)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", errors.New("could not generate token")
	}

	return tokenString, nil
}

func getToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value")
	}

	bearer_token := strings.Split(header, " ")

	if len(bearer_token) != 2 {
		return "", errors.New("bad header format")
	}

	return bearer_token[1], nil
}

func parseToken(jwtToken string, api_key string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		_, OK := token.Method.(*jwt.SigningMethodHMAC)
		if !OK {
			return nil, errors.New("bad signing method")
		}
		return []byte(api_key), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func checkTokenExpiration(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		expiration := time.Unix(int64(claims["expiration"].(float64)), 0)
		if expiration.Before(time.Now()) {
			return errors.New("token expired")
		}
	} else {
		return errors.New("bad token claims")
	}

	return nil
}

func IsAuthCheck(c *gin.Context, api_key string) {
	auth_header := c.GetHeader("Authorization")

	token, err := getToken(auth_header)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	jwt_token, err := parseToken(token, api_key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = checkTokenExpiration(jwt_token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Next()
}
