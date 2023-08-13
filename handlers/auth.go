package handlers

import (
	"net/http"
	"time"

	"github.com/awalvie/recall/auth"
	"github.com/awalvie/recall/config"
	"github.com/labstack/echo/v4"
)

// Login is the handler for the login route
// it checks if the username and password are correct
// and sets the session cookie if they are
func Login(e echo.Context) error {
	// Get the config from the context
	a := e.Get("app").(*config.App)

	// Get username and password from the request
	username := e.FormValue("username")
	password := e.FormValue("password")

	// Check if the username and password are correct
	if username == a.Config.Auth.Username && password == a.Config.Auth.Password {
		// Set auth session cookie
		cookie := &http.Cookie{
			Name:    "auth",
			Value:   username + ":" + auth.Secret(username, password),
			Expires: time.Now().Add(time.Hour * 24 * 7), // 1 week
		}
		e.SetCookie(cookie)

		// Redirect to the index page
		return e.Redirect(http.StatusFound, "/")
	}

	// Incorrect credentials, redirect to the login page with error message
	return e.Redirect(http.StatusFound, "/login?error=1")
}

// Logout is the handler for the logout route
// it deletes the session cookie and redirects to the index page
func Logout(e echo.Context) error {
	// Delete the session cookie
	cookie := &http.Cookie{
		Name:    "auth",
		Value:   "",
		Expires: time.Now().Add(time.Hour * -1),
	}
	e.SetCookie(cookie)

	// Redirect to the index page
	return e.Redirect(http.StatusFound, "/")
}
