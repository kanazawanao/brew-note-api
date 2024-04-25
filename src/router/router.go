package router

import (
	"brew-note/src/handler"

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
	e.GET("/me", handler.GetUsers)

	e.GET("/users", handler.GetUsers)
	e.POST("/users", handler.PostUsers)
	e.GET("/users/:id", handler.GetUser)

	e.POST("/beans", handler.PostBean)
	e.GET("/beans", handler.GetBeans)
	e.GET("/beans/:beanId", handler.GetBean)

	return e
}
