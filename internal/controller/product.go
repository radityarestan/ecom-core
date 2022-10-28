package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/radityarestan/ecom-core/internal/service"
	"github.com/radityarestan/ecom-core/internal/shared/dto"
	"go.uber.org/dig"
	"net/http"
	"strconv"
)

type Product struct {
	dig.In
	Service service.Holder
}

func (impl *Product) Create(c echo.Context) error {
	var (
		ctx = context.WithValue(c.Request().Context(), "user_id", c.Get("user_id"))
		req = dto.CreateProductRequest{}
	)

	if err := bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	res, err := impl.Service.Product.CreateProduct(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	if err := c.Validate(res); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  dto.StatusSuccess,
		Message: dto.CreateProductSuccess,
		Data:    res,
	})
}

func (impl *Product) Catalog(c echo.Context) error {
	var (
		ctx    = c.Request().Context()
		userID = c.Get("user_id").(uint)
	)

	res, err := impl.Service.Product.GetBaseProducts(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	if err := c.Validate(res); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  dto.StatusSuccess,
		Message: dto.CatalogProductSuccess,
		Data:    res,
	})
}

func (impl *Product) Search(c echo.Context) error {
	var ctx = c.Request().Context()
	search := c.QueryParam("search")

	if search == "" {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: dto.SearchProductError,
		})
	}

	res, err := impl.Service.Product.FindProducts(ctx, search)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	if err := c.Validate(res); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  dto.StatusSuccess,
		Message: dto.SearchProductSuccess,
		Data:    res,
	})
}

func (impl *Product) Detail(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = c.Param("id")
	)

	id, err := strconv.Atoi(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	res, err := impl.Service.Product.FindProductByID(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	if err := c.Validate(res); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  dto.StatusSuccess,
		Message: dto.DetailProductSuccess,
		Data:    res,
	})
}
