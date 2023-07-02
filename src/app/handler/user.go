package handler

import (
	"fmt"
	"net/http"
	"tripig/src/app/database"
	"tripig/src/app/models"

	"github.com/labstack/echo/v4"
)

type Users []models.User

// e.Get("/users/", GetUsers)
func GetUsers(c echo.Context) error {
	var users Users
	result := database.Handler.Find(&users)

	if err := result.Error; err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// e.POST("/users/", PostUserss)
func PostUsers(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	database.Handler.Create(&user)
	return c.JSON(http.StatusOK, user)
}
