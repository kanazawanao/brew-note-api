package router

import (
	"brew-note/src/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"Authorization"},
	}))
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

	e.GET("roast-levels", handler.GetRoastLevels)

	return e
}
