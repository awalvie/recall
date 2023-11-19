package handlers

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/awalvie/recall/config"
	"github.com/awalvie/recall/models"
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

func ContactsPage(e echo.Context) error {
	// Get the config from the context
	a := e.Get("app").(*config.App)

	// Get cookie from the context
	cookie, err := e.Cookie("auth")
	if err != nil {
		log.Println("error getting cookie:", err)
		return err
	}

	// Bulid url for /api/contacts
	url := "http://" + a.Config.Server.Host + ":" + strconv.Itoa(a.Config.Server.Port) + "/api/contacts"

	// Create a new request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error creating request:", err)
		return err
	}

	// Set the cookie in the request
	req.AddCookie(cookie)

	// Make the http request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error getting contacts:", err)
		return err
	}

	// Read the Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response:", err)
		return err
	}

	// Unmarshal body into a slice of contacts
	var contacts []models.Contact
	err = json.Unmarshal(body, &contacts)
	if err != nil {
		log.Println("error unmarshaling response:", err)
		return err
	}

	// Get the template path
	tpath := filepath.Join(a.Config.Dirs.Templates, "*")

	// Parse the Templates
	t := template.Must(template.ParseGlob(tpath))

	// Render the Templates
	if err := t.ExecuteTemplate(e.Response().Writer, "contacts", contacts); err != nil {
		log.Println("error rendering template:", err)
		return err
	}

	return nil
}

// LoginPage is the handler for the login route
// that renders the login template
func LoginPage(e echo.Context) error {
	// Get the config from the context
	a := e.Get("app").(*config.App)

	// Get the error message from query paramater
	loginError := e.QueryParam("error")

	// Check if the error message is set
	data := map[string]string{}
	if loginError == "1" {
		data["error"] = "Incorrect username or password"
	}

	// Get the template path
	tpath := filepath.Join(a.Config.Dirs.Templates, "*")

	// Parse the templates
	t := template.Must(template.ParseGlob(tpath))

	// Render the templates
	if err := t.ExecuteTemplate(e.Response().Writer, "login", data); err != nil {
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

// DashboardPage redirects to /
func DashboardPage(e echo.Context) error {
	e.Redirect(http.StatusFound, "/")
	return nil
}

// PageNotFound renders a 404 page when a url that does
// not exist is accessed
func PageNotFound(err error, e echo.Context) {
	// Get app config from the Context
	a := e.Get("app").(*config.App)

	// Get the template path
	tpath := filepath.Join(a.Config.Dirs.Templates, "*")

	// Parse the templates
	t := template.Must(template.ParseGlob(tpath))

	// Render the templates
	if err := t.ExecuteTemplate(e.Response().Writer, "404", nil); err != nil {
		log.Println("error rendering template:", err)
	}
}

// AddContactPage renders a page for adding a contact
// with a form that make a POST call to /api/contacts and
// adds a contact to the database
func AddContactPage(e echo.Context) error {
	// Get app config from the Context
	a := e.Get("app").(*config.App)

	// Get the template path
	tpath := filepath.Join(a.Config.Dirs.Templates, "*")

	// Parse the templates
	t := template.Must(template.ParseGlob(tpath))

	// Render the templates
	if err := t.ExecuteTemplate(e.Response().Writer, "add-contact", nil); err != nil {
		log.Println("error rendering template:", err)
	}

	return nil
}

// ContactPage renders a single page showing all the details
// about an existing contact
func ContactPage(e echo.Context) error {
	// Get app config from the Context
	a := e.Get("app").(*config.App)

	// Get the contact id from the url
	id := e.Param("id")

	// Get cookie from the context
	cookie, err := e.Cookie("auth")
	if err != nil {
		log.Println("error getting cookie:", err)
		return err
	}

	// Bulid url for /api/contacts
	url := "http://" + a.Config.Server.Host + ":" + strconv.Itoa(a.Config.Server.Port) + "/api/contacts/" + id

	// Create a new request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error creating request:", err)
		return err
	}

	// Set the cookie in the request
	req.AddCookie(cookie)

	// Make the http request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error getting contacts:", err)
		return err
	}

	// Read the Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response:", err)
		return err
	}

	// Unmarshal body into a contact
	var contact models.Contact
	err = json.Unmarshal(body, &contact)
	if err != nil {
		log.Println("error unmarshaling response:", err)
		return err
	}

	data := struct {
		Contact models.Contact
	}{
		Contact: contact,
	}

	// Get the template path
	tpath := filepath.Join(a.Config.Dirs.Templates, "*")

	// Parse the templates
	t := template.Must(template.ParseGlob(tpath))

	// Render the templates
	if err := t.ExecuteTemplate(e.Response().Writer, "contact", data); err != nil {
		log.Println("error rendering template:", err)
	}
	return nil
}
