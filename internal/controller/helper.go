package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/radityarestan/ecom-core/internal/service"
)

const (
	BearerLength = 7
)

func bind(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	return nil
}

func getUserIDFromJWT(c echo.Context) uint {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return 0
	}
	token = token[BearerLength:]

	claims := &service.Claims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return service.JWTKey, nil
	})
	if err != nil {
		return 0
	}

	if !jwtToken.Valid {
		return 0
	}

	return claims.ID
}
