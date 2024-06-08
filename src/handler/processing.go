package handler

import (
	"brew-note/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProcessings(c echo.Context) error {
	processings := services.GetProcessings()
	return c.JSON(http.StatusOK, processings)
}
