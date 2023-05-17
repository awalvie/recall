package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/awalvie/recall/config"
	"github.com/labstack/echo/v4"
)

// Index is the handler for the index route
func Index(e echo.Context) error {
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

// Static is the handler for serving all static files
func Static(e echo.Context) error {
	// Get app config from the Context
	a := e.Get("app").(*config.App)

	// Get file from the context
	f := e.Param("file")
	f = filepath.Clean(filepath.Join(a.Config.Dirs.Static, f))

	http.ServeFile(e.Response().Writer, e.Request(), f)
	log.Println("served file:", f)

	return nil
}
