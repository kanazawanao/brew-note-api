package handler

import (
	"coffee-paws/src/services"
	"net/http"

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

func GetPlaceTypes(c echo.Context) error {
	res := services.GetPlaceTypeList();
	return c.JSON(http.StatusOK, res)
}
