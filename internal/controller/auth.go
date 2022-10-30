package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/radityarestan/ecom-core/internal/service"
	"github.com/radityarestan/ecom-core/internal/shared/dto"
	"go.uber.org/dig"
	"net/http"
)

type Auth struct {
	dig.In
	Service service.Holder
}

func (impl *Auth) SignIn(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = dto.SignInRequest{}
	)

	if err := bind(c, &req); err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	res, err := impl.Service.Auth.SignIn(ctx, &req)
	if err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	if err := c.Validate(res); err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  dto.StatusSuccess,
		Message: dto.SignInSuccess,
		Data:    res,
	})
}

func (impl *Auth) SignUp(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = dto.SignUpRequest{}
	)

	if err := bind(c, &req); err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	res, err := impl.Service.Auth.SignUp(ctx, &req)
	if err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	if err := c.Validate(res); err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Status:  dto.StatusSuccess,
		Message: dto.CreateUserSuccess,
		Data:    res,
	})
}

func (impl *Auth) Verify(c echo.Context) error {
	var (
		ctx  = c.Request().Context()
		code = c.Param("code")
	)

	if code == "" {
		c.Set(dto.StatusError, dto.ErrInvalidCode.Error())
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: dto.ErrInvalidCode.Error(),
		})
	}

	res, err := impl.Service.Auth.VerifyEmail(ctx, code)
	if err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	if err := c.Validate(res); err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  dto.StatusSuccess,
		Message: dto.VerifyEmailSuccess,
		Data:    res,
	})
}
