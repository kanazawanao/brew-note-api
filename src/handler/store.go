package handler

import (
	"coffee-paws/src/models"
	"coffee-paws/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.POST("/stores/", PostStore)
func PostStore(c echo.Context) error {
	store := models.Store{}
	res := services.PostStore(store)
	return c.JSON(http.StatusOK, res)
}