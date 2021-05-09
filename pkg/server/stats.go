package server

import (
	"net/http"

	"github.com/gulien/fizz-buzz/pkg/stats"
	"github.com/labstack/echo/v4"
)

func statsHandler(statistics stats.Statistics) echo.HandlerFunc {
	return func(c echo.Context) error {
		entry, err := statistics.GetMostFrequentEntry()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, entry)
	}
}
