package handlers

import "github.com/labstack/echo/v4"

func Index(e echo.Context) error {
	return e.String(200, "Hello, world!")
}
