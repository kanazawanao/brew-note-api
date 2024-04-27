package handler

import (
	"brew-note/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.Get("/roast-levels/", GetRoastLevels)
func GetRoastLevels(c echo.Context) error {
	// token := c.Request().Header.Get("Brew-Note-Auth-Token")
	users := services.GetRoastLevels()

	return c.JSON(http.StatusOK, users)
}
