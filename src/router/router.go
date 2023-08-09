package router

import (
	"coffee-paws/src/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// healthcheck
	e.GET("/healthcheck", handler.Healthcheck)

	e.GET("/coffee-paws/users", handler.GetUsers)
	e.POST("/coffee-paws/users", handler.PostUsers)
	e.GET("/coffee-paws/users/:id", handler.GetUSer)

	e.GET("/coffee-paws/places/nearby", handler.GetNearbySearch)
	e.GET("/coffee-paws/places/types", handler.GetPlaceTypes)

	e.POST("/coffee-paws/stores", handler.PostStore)

	return e
}
