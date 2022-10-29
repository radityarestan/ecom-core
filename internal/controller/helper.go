package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"strconv"
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

func checkIntParam(param ...string) ([]int, error) {
	var res = make([]int, len(param))

	for i, p := range param {
		if p == "" {
			return nil, errors.New("limit offset param cannot be empty")
		}

		val, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}

		res[i] = val
	}

	return res, nil
}
