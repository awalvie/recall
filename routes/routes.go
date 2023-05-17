package routes

import (
	"log"

	"github.com/awalvie/recall/config"
	"github.com/awalvie/recall/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Configure configures the echo server
func Configure(e *echo.Echo, c *config.Config) {
	// Hide the stupid banner
	e.HideBanner = true
	e.HidePort = true

	// Configure middleware

	// Pass app config to handlers
	e.Use(ConfigMiddleware(*c))

	// Log requests
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

	// Recover from panics
	e.Use(middleware.Recover())

	// Configure routes
	e.GET("/", handlers.Index)
}

// ConfigMiddleware adds the config to the context
func ConfigMiddleware(config config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Add the config to the context
			c.Set("config", config)

			// Call the next handler in the chain
			return next(c)
		}
	}
}
