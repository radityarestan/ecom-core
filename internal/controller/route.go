package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radityarestan/ecom-authentication/internal/shared"
	"go.uber.org/dig"
)

const PrefixAPI = "/api/auth"

type CustomValidator struct {
	validator *validator.Validate
}

type Holder struct {
	dig.In
	Deps shared.Deps
	Auth Auth
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (h *Holder) RegisterRoutes() {
	var app = h.Deps.Server

	app.Validator = &CustomValidator{validator: validator.New()}
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	app.POST(PrefixAPI, h.Auth.Post)
}
