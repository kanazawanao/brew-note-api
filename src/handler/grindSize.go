package handler

import (
	"brew-note/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetGrindSizes(c echo.Context) error {
	grindSizes := services.GetGrindSizes()
	return c.JSON(http.StatusOK, grindSizes)
}
