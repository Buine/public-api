package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/tesis/user-ms/config"
	"github.com/tesis/user-ms/errors"
	"github.com/tesis/user-ms/responses"
	"strings"
	"time"
)

type CustomClaim struct {
	responses.UserResponse
	jwt.StandardClaims
}

var mySigningKey = []byte(config.ServerConfig{}.JwtKey)

func ValidateJwt(context echo.Context) (*responses.UserResponse, error) {
	tokenString := strings.ReplaceAll(context.Request().Header.Get("Authorization"), "Bearer ", "")
	if tokenString == "" {
		return nil, &errors.Unauthenticated
	}
	token, err := ReaderToken(tokenString)
	if err != nil {
		return nil, err
	}

	context.Request().Header.Set("x-user-code", token.Code)

	return token, nil
}

func BuildToken(userData responses.UserResponse) (string, error) {

	claims := CustomClaim{
		userData,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			Issuer:    "QueryBuilderTesis",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ReaderToken(tokenString string) (*responses.UserResponse, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		return &claims.UserResponse, nil
	} else {
		return nil, &errors.SessionExpired
	}
}
