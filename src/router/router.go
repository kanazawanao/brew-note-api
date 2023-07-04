package router

import (
	"tripig/src/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// healthcheck
	e.GET("/healthcheck", handler.Healthcheck)

	e.GET("/tripig/users", handler.GetUsers)
	e.POST("/tripig/users", handler.PostUsers)

	return e
}
