package handlers

import (
	"net/http"
	"strconv"

	"github.com/awalvie/recall/config"
	"github.com/awalvie/recall/models"
	"github.com/labstack/echo/v4"
)

// GetContacts returns all contacts
func GetContacts(e echo.Context) error {
	a := e.Get("app").(*config.App)
	db := a.DB

	// Get all contacts
	var contacts []models.Contact
	db.Find(&contacts)
	return e.JSON(http.StatusOK, contacts)
}

// GetContact returns a single contact
func GetContact(e echo.Context) error {
	a := e.Get("app").(*config.App)
	db := a.DB

	// Get the ID from the URL
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid ID")
	}

	// Get the contact
	var contact models.Contact
	result := db.First(&contact, id)
	if result.Error != nil {
		return e.String(http.StatusNotFound, "Contact not found")
	}

	return e.JSON(http.StatusOK, contact)
}

// CreateContact creates a new contact
func CreateContact(e echo.Context) error {
	// Get the config from the context
	a := e.Get("app").(*config.App)
	db := a.DB

	contact := new(models.Contact)
	if err := e.Bind(contact); err != nil {
		return e.String(http.StatusBadRequest, "Invalid data")
	}

	db.Create(&contact)
	return e.JSON(http.StatusCreated, contact)
}

// UpdateContact updates a contact
func UpdateContact(e echo.Context) error {
	// Get the config from the context
	a := e.Get("app").(*config.App)
	db := a.DB

	// Get the ID from the URL
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid ID")
	}

	// Get the contact from the database
	contact := new(models.Contact)
	result := db.First(&contact, id)
	if result.Error != nil {
		return e.String(http.StatusNotFound, "Contact not found")
	}

	if err := e.Bind(contact); err != nil {
		return e.String(http.StatusBadRequest, "Invalid data")
	}

	db.Save(&contact)
	return e.JSON(http.StatusOK, contact)
}

// DeleteContact deletes a contact
func DeleteContact(e echo.Context) error {
	// Get the config from the context
	a := e.Get("app").(*config.App)
	db := a.DB

	// Get the ID from the URL
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid ID")
	}

	// Get the contact from the database
	contact := new(models.Contact)
	result := db.First(&contact, id)
	if result.Error != nil {
		return e.String(http.StatusNotFound, "Contact not found")
	}

	db.Delete(&contact)
	return e.NoContent(http.StatusNoContent)
}
