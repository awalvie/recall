package routes

import (
	"log"
	"net/http"

	"github.com/awalvie/recall/auth"
	"github.com/awalvie/recall/config"
	"github.com/awalvie/recall/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Configure sets up the routes and middlwares for the application
func Configure(e *echo.Echo, a *config.App) {
	// Hide the stupid banner
	e.HideBanner = true
	e.HidePort = true

	// Middleware configuration

	// Recover from panics
	e.Use(middleware.Recover())

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

	// Pass app config to handlers
	e.Use(AppConfigMiddleware(a))

	// Add auth middleware
	e.Use(AuthMiddlware(a))

	// Add custom error handler
	e.HTTPErrorHandler = handlers.PageNotFound

	// Configure routes
	// Page Routes
	e.GET("/", handlers.IndexPage)
	e.GET("/contacts", handlers.ContactsPage)
	e.GET("/static/:file", handlers.StaticFiles)
	e.GET("/dashboard", handlers.DashboardPage)
	e.GET("/add-contact", handlers.AddContactPage)

	// Auth routes
	e.GET("/login", handlers.LoginPage)
	e.GET("/logout", handlers.Logout)
	e.POST("/login", handlers.Login)

	// API routes
	api := e.Group("/api")
	api.GET("/contacts", handlers.GetContacts)
	api.POST("/contacts", handlers.CreateContact)
	api.GET("/contacts/:id", handlers.GetContact)
	api.PUT("/contacts/:id", handlers.UpdateContact)
	api.DELETE("/contacts/:id", handlers.DeleteContact)
}

// AppConfigMiddleware adds the config to the context
func AppConfigMiddleware(app *config.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Add the config to the context
			c.Set("app", app)

			// Call the next handler in the chain
			return next(c)
		}
	}
}

// AuthMiddleware does the following:
// 1. Checks if the request is being made to a public page
// 2. If not, check if the user is authenticated
// 3. If not, redirect to the login page
func AuthMiddlware(app *config.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// Check if request is being made to a public page
			if auth.IsPublic(c.Path()) {
				log.Println("public route: ", c.Path())
				return next(c)
			}

			// Check if the user is authenticated
			if !auth.IsAuthenticated(c) {
				// Redirect to the login page
				return c.Redirect(http.StatusFound, "/login")
			}

			// Call the next handler in the chain
			return next(c)
		}
	}
}
