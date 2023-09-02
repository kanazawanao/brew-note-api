package handler

import (
	"coffee-paws/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.Get("/users/", GetUsers)
func GetProfile(c echo.Context) error {
	token := c.Request().Header.Get("Coffee-Paws-Auth-Token")
	users := services.GetProfile(token)

	return c.JSON(http.StatusOK, users)
}