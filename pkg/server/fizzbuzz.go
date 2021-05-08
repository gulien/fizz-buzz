package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.come/gulien/fizz-buzz/pkg/fizzbuzz"
)

func fizzBuzzHandler(timeout time.Duration) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		var (
			int1, int2, limit int
			str1, str2        string
		)

		// err = first error.
		err := echo.QueryParamsBinder(c).
			FailFast(true).
			MustInt("int1", &int1).
			MustInt("int2", &int2).
			MustInt("limit", &limit).
			// We accept empty str1/str2.
			String("str1", &str1).
			String("str2", &str2).
			BindError()

		if err != nil {
			errBinding, ok := err.(*echo.BindingError)

			if !ok {
				// Unexpected error kind: fail loud.
				return err
			}

			// Format the error so it's more in line with other errors.
			return echo.NewHTTPError(
				http.StatusBadRequest,
				fmt.Errorf("%s %s", errBinding.Field, errBinding.Message).Error(),
			)
		}

		result, err := fizzbuzz.FizzBuzz(ctx, int1, int2, limit, str1, str2)

		if err == nil {
			return c.JSON(http.StatusOK, result)
		}

		if errors.Is(err, fizzbuzz.ErrZeroInt) || errors.Is(err, fizzbuzz.ErrNegativeOrZeroLimit) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if errors.Is(err, ctx.Err()) {
			return echo.NewHTTPError(http.StatusServiceUnavailable, err.Error())
		}

		// Unexpected error: fail loud.
		return err
	}
}
