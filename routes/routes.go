package routes

import (
	"log"

	"github.com/awalvie/recall/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Configure configures the echo server
func Configure(e *echo.Echo) {
	// Hide the stupid banner
	e.HideBanner = true

	// Configure middleware
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogError:   true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Println("request",
				"URI:", v.URI,
				"status:", v.Status,
				"error", v.Error,
				"method", v.Method,
				"latency", v.Latency,
			)
			return nil
		},
	}))
	e.Use(middleware.Recover())

	// Configure routes
	e.GET("/", handlers.Index)
}
