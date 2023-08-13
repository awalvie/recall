package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/awalvie/recall/config"
	"github.com/labstack/echo/v4"
)

// IndexPage is the handler for the index route
// that renders the index template
func IndexPage(e echo.Context) error {
	// Get the config from the context
	a := e.Get("app").(*config.App)

	// Get the template path
	tpath := filepath.Join(a.Config.Dirs.Templates, "*")

	// Parse the templates
	t := template.Must(template.ParseGlob(tpath))

	// Render the templates
	if err := t.ExecuteTemplate(e.Response().Writer, "index", nil); err != nil {
		log.Println("error rendering template:", err)
		return err
	}
	return nil
}

// StaticFiles is the handler for serving all static files
// from the configured static directory
func StaticFiles(e echo.Context) error {
	// Get app config from the Context
	a := e.Get("app").(*config.App)

	// Get file from the context
	f := e.Param("file")
	f = filepath.Clean(filepath.Join(a.Config.Dirs.Static, f))

	http.ServeFile(e.Response().Writer, e.Request(), f)
	log.Println("served file:", f)

	return nil
}

// Login is the handler for serving the login page
func LoginPage(e echo.Context) error {
	// Get app config from the context
	a := e.Get("app").(*config.App)

	// Get the template path
	tpath := filepath.Join(a.Config.Dirs.Templates, "*")

	// Parse the Templates
	t := template.Must(template.ParseGlob(tpath))

	// Render the Templates
	if err := t.ExecuteTemplate(e.Response().Writer, "login", nil); err != nil {
		log.Println("error rendering template:", err)
		return err
	}
	return nil

}
