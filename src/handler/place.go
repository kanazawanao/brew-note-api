package handler

import (
	"net/http"
	"tripig/src/services"

	"github.com/labstack/echo/v4"
)

// e.Get("/place/search/nearby", GetNearbySearch)
func GetNearbySearch(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	pageToken := c.QueryParam("pageToken")
	placeType := c.QueryParam("placeType")
	res := services.NearbySearch(keyword, pageToken, placeType)
	return c.JSON(http.StatusOK, res)
}
