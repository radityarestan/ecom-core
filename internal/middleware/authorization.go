package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/radityarestan/ecom-core/internal/service"
	"github.com/radityarestan/ecom-core/internal/shared/dto"
	"net/http"
)

const (
	BearerLength = 7
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			c.Set(dto.StatusError, dto.Unauthorized)
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Status:  dto.StatusError,
				Message: dto.Unauthorized,
			})
		}
		token = token[BearerLength:]
		claims := &service.Claims{}
		jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return service.JWTKey, nil
		})

		if err != nil {
			c.Set(dto.StatusError, dto.Unauthorized)
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Status:  dto.StatusError,
				Message: dto.Unauthorized,
			})
		}

		if !jwtToken.Valid {
			c.Set(dto.StatusError, dto.Unauthorized)
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Status:  dto.StatusError,
				Message: dto.Unauthorized,
			})
		}

		c.Set("user_id", claims.ID)

		return next(c)
	}
}
