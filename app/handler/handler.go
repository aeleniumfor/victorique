package handler

import (
	"github.com/labstack/echo"
)

// GetUserContainers is Containerlist
func GetUserContainers() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "HelloWorld")
	}
}
