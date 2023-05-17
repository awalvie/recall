package routes

import (
	"log"

	"github.com/awalvie/recall/config"
	"github.com/awalvie/recall/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Configure configures the echo server
func Configure(e *echo.Echo, a *config.App) {
	// Hide the stupid banner
	e.HideBanner = true
	e.HidePort = true

	// Middleware configuration

	// Pass app config to handlers
	e.Use(ConfigMiddleware(a))

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
	e.GET("/static/:file", handlers.Static)
}

// ConfigMiddleware adds the config to the context
func ConfigMiddleware(app *config.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Add the config to the context
			c.Set("app", app)

			// Call the next handler in the chain
			return next(c)
		}
	}
}
