package handler

import (
	"coffee-paws/src/models"
	"coffee-paws/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.Get("/users/", GetUsers)
func GetUsers(c echo.Context) error {
	users := services.GetUsers()

	return c.JSON(http.StatusOK, users)
}

// e.POST("/users/", PostUserss)
func PostUsers(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// e.Get("/users/{id}", GetUser)
func GetUser(c echo.Context) error {
	id := c.Param("id")
	user := services.GetUser(id)
	return c.JSON(http.StatusOK, user)
}
