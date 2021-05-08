package server

import (
	"github.com/labstack/echo/v4"
)

func pingHandler() echo.HandlerFunc {
	return func(_ echo.Context) error {
		return nil
	}
}
