package server

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// New returns an instance of echo.Echo.
func New(timeout time.Duration) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// Middlewares.
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			// No logging for the ping handler as it may output a lot of logs
			// according to the health check strategy.
			return c.Request().URL.Path == "/"
		},
	}))
	e.Use(middleware.Recover())

	// Routes.
	e.GET("/", pingHandler())
	api := e.Group("/api/v1")
	api.GET("/fizz-buzz", fizzBuzzHandler(timeout))

	return e
}
