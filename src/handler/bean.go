package handler

import (
	"brew-note/src/models"
	"brew-note/src/services"
	"brew-note/src/utils"
	"log"
	"net/http"

	openapi "github.com/brew-note/api"

	"github.com/labstack/echo/v4"
)

// e.POST("/beans", PostBean)
func PostBean(c echo.Context) error {
	s := new(openapi.CreateBean)
	if err := c.Bind(s); err != nil {
		log.Printf("err %v", err.Error())
		return c.String(http.StatusBadRequest, "bad request")
	}
	id := utils.GenerateULID()
	bean := models.Bean{
		ID: id,
		UserId: "",
		ProductionArea: s.ProductionArea,
		PlantationName: s.PlantationName,
		Kind: s.Kind,
		RoastLevel: s.RoastLevel,
	}

	res := services.PostBean(bean)
	return c.JSON(http.StatusOK, res)
}

// e.Get("/beans", GetBeans)
func GetBeans(c echo.Context) error {
	storeId := c.Param("storeId")
	beans := services.GetBeans(storeId)

	return c.JSON(http.StatusOK, beans)
}

// e.Get("/beans/:beanId", GetBean)
func GetBean(c echo.Context) error {
	beanId := c.Param("beanId")
	bean := services.GetBean(beanId)

	return c.JSON(http.StatusOK, bean)
}