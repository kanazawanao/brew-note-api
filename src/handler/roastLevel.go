package handler

import (
	"brew-note/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRoastLevels(c echo.Context) error {
	roastLevels := services.GetRoastLevels()
	return c.JSON(http.StatusOK, roastLevels)
}
