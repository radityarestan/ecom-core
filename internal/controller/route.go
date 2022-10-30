package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	customMiddleware "github.com/radityarestan/ecom-core/internal/middleware"
	"github.com/radityarestan/ecom-core/internal/shared"
	"go.uber.org/dig"
	"unicode"
	"unicode/utf8"
)

const (
	PrefixAuthAPI    = "/api/auth"
	PrefixProductAPI = "/api/product"

	SignUpAPI = "/sign-up"
	SignInAPI = "/sign-in"
	VerifyAPI = "/verify/:code"

	ProductSearchAPI = "/search"
	ProductDetailAPI = "/:id"
)

type CustomValidator struct {
	validator *validator.Validate
}

type Holder struct {
	dig.In
	Deps    shared.Deps
	Auth    Auth
	Product Product
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (h *Holder) RegisterRoutes() {
	var app = h.Deps.Server

	newValidator := initValidator()
	app.Validator = &CustomValidator{validator: newValidator}

	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(customMiddleware.MetricsMiddleware)

	app.GET("/prometheus", echo.WrapHandler(promhttp.Handler()))

	authRoutes := app.Group(PrefixAuthAPI)
	{
		authRoutes.POST(SignUpAPI, h.Auth.SignUp)
		authRoutes.POST(SignInAPI, h.Auth.SignIn)
		authRoutes.GET(VerifyAPI, h.Auth.Verify)
	}

	productRoutes := app.Group(PrefixProductAPI)
	productRoutes.Use(customMiddleware.AuthMiddleware)
	{
		productRoutes.POST("", h.Product.Create)
		productRoutes.GET("", h.Product.Catalog)
		productRoutes.GET(ProductSearchAPI, h.Product.Search)
		productRoutes.GET(ProductDetailAPI, h.Product.Detail)
	}

}

func initValidator() *validator.Validate {
	v := validator.New()
	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		var (
			hasNumber      = false
			hasSpecialChar = false
			hasLetter      = false
			hasSuitableLen = false
		)

		password := fl.Field().String()

		if utf8.RuneCountInString(password) <= 30 || utf8.RuneCountInString(password) >= 6 {
			hasSuitableLen = true
		}

		for _, c := range password {
			switch {
			case unicode.IsNumber(c):
				hasNumber = true
			case unicode.IsPunct(c) || unicode.IsSymbol(c):
				hasSpecialChar = true
			case unicode.IsLetter(c) || c == ' ':
				hasLetter = true
			default:
				return false
			}
		}

		return hasNumber && hasSpecialChar && hasLetter && hasSuitableLen
	})

	return v
}
