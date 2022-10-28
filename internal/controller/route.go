package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radityarestan/ecom-core/internal/shared"
	"go.uber.org/dig"
	"unicode"
	"unicode/utf8"
)

const (
	PrefixAuthAPI    = "/api/auth"
	PrefixProductAPI = "/api/product"

	SignUpAPI = PrefixAuthAPI + "/sign-up"
	SignInAPI = PrefixAuthAPI + "/sign-in"
	VerifyAPI = PrefixAuthAPI + "/verify/:code"

	ProductAPI       = PrefixProductAPI
	ProductSearchAPI = PrefixProductAPI + "/search"
	ProductDetailAPI = PrefixProductAPI + "/:id"
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

	app.POST(SignUpAPI, h.Auth.SignUp)
	app.POST(SignInAPI, h.Auth.SignIn)
	app.GET(VerifyAPI, h.Auth.Verify)

	app.POST(ProductAPI, h.Product.Create)
	app.GET(ProductAPI, h.Product.Catalog)
	app.GET(ProductSearchAPI, h.Product.Search)
	app.GET(ProductDetailAPI, h.Product.Detail)
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
