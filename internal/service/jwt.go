package service

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	ID uint
	jwt.StandardClaims
}

var JWTKey = []byte("my_secret_key")

func (a *authService) generateToken(userID uint) (string, error) {
	expTime := time.Now().Add(time.Hour * 24).Unix()

	claims := &Claims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}
