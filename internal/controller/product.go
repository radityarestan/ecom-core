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
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	res, err := impl.Service.Product.CreateProduct(ctx, &req)
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
		Message: dto.CreateProductSuccess,
		Data:    res,
	})
}

func (impl *Product) UploadPhoto(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = c.FormValue("id")
	)

	params, err := checkIntParam(id)
	if err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	f, err := c.FormFile("file")
	if err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	// check file size
	if f.Size > 2*1024*1024 {
		c.Set(dto.StatusError, dto.FileSizeExceeded)
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: dto.FileSizeExceeded,
		})
	}

	if err := impl.Service.Product.UploadProductPhoto(ctx, f, uint(params[0])); err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Status:  dto.StatusSuccess,
		Message: dto.UploadProductPhotoSuccess,
	})
}

func (impl *Product) Catalog(c echo.Context) error {
	var (
		ctx    = c.Request().Context()
		userID = c.Get("user_id")
		limit  = c.QueryParam("lmt")
		offset = c.QueryParam("oft")
	)

	paramInt, err := checkIntParam(limit, offset)
	if err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	res, err := impl.Service.Product.GetProductCatalog(ctx, userID.(uint), paramInt[0], paramInt[1])
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
		Message: dto.CatalogProductSuccess,
		Data:    res,
	})
}

func (impl *Product) Search(c echo.Context) error {
	var (
		ctx    = c.Request().Context()
		search = c.QueryParam("q")
		limit  = c.QueryParam("lmt")
		offset = c.QueryParam("oft")
	)

	paramInt, err := checkIntParam(limit, offset)
	if err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	res, err := impl.Service.Product.FindProducts(ctx, search, paramInt[0], paramInt[1])
	if err != nil {
		c.Set(dto.StatusError, err.Error())
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
		ctx = context.WithValue(c.Request().Context(), "user_id", c.Get("user_id"))
		req = c.Param("id")
	)

	id, err := strconv.Atoi(req)
	if err != nil {
		c.Set(dto.StatusError, err.Error())
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	res, err := impl.Service.Product.FindProductByID(ctx, uint(id))
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
		Message: dto.DetailProductSuccess,
		Data:    res,
	})
}
