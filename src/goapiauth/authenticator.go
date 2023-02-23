package goapiauth

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
)

type (
	Authy interface {
		GetUserFromToken(string) (string, error)
		GenerateToken(string) string
	}

	Authenticator struct {
		Config *AuthConfig
	}
)

func New() (auth *Authenticator, err error) {
	auth = &Authenticator{}
	if err = auth.readConfig(); err != nil {
		return
	}
	return
}

func (auth *Authenticator) GenerateToken(userId string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
	})
	tokenString, err = token.SignedString([]byte(auth.Config.SigningKey))
	return
}

func (auth *Authenticator) GetUserFromToken(tokenString string) (userId string, err error) {
	var (
		token *jwt.Token
	)
	if token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(auth.Config.SigningKey), nil
	}); err != nil {
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId = claims["userId"].(string)
		return
	}
	return
}
