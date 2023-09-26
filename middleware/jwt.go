package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)


type jwtCustomClaims struct {
	UserId  int `json:"userId"`
	Name string   `json:"bool"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userId int, name string) string {
	var claims = jwtCustomClaims {
		userId,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resultJWT, _ := token.SignedString([]byte("123"))
	return resultJWT
}