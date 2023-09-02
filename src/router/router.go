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
	e.GET("/coffee-paws/me", handler.GetUsers)

	e.GET("/coffee-paws/users", handler.GetUsers)
	e.POST("/coffee-paws/users", handler.PostUsers)
	e.GET("/coffee-paws/users/:id", handler.GetUser)

	e.GET("/coffee-paws/places/nearby", handler.GetNearbySearch)
	e.GET("/coffee-paws/places/types", handler.GetPlaceTypes)

	e.POST("/coffee-paws/stores", handler.PostStore)
	e.GET("/coffee-paws/stores", handler.GetStores)
	e.GET("/coffee-paws/stores/:id", handler.GetStore)

	e.POST("/coffee-paws/stores/:storeId/beans", handler.PostBean)
	e.GET("/coffee-paws/stores/:storeId/beans", handler.GetBeans)
	e.GET("/coffee-paws/stores/:storeId/beans/:beanId", handler.GetBean)

	return e
}
