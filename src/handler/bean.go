package handler

import (
	"coffee-paws/src/models"
	"coffee-paws/src/services"
	"coffee-paws/src/utils"
	"log"
	"net/http"

	openapi "github.com/coffee-paws/api"

	"github.com/labstack/echo/v4"
)

// e.POST("stores/:storeId/beans", PostBean)
func PostBean(c echo.Context) error {
	storeId := c.Param("storeId")
	s := new(openapi.CreateBean)
	if err := c.Bind(s); err != nil {
		log.Printf("err %v", err.Error())
		return c.String(http.StatusBadRequest, "bad request")
	}
	id := utils.GenerateULID()
	bean := models.Bean{
		ID: id,
		StoreId: storeId,
		ProductionArea: s.ProductionArea,
		PlantationName: s.PlantationName,
		Kind: s.Kind,
		RoastLevel: s.RoastLevel,
		Price: s.Price,
	}

	res := services.PostBean(bean)
	return c.JSON(http.StatusOK, res)
}

// e.Get("stores/:storeId/beans", GetBeans)
func GetBeans(c echo.Context) error {
	storeId := c.Param("storeId")
	beans := services.GetBeans(storeId)

	return c.JSON(http.StatusOK, beans)
}

// e.Get("/stores/:storeId/beans/:beanId", GetBean)
func GetBean(c echo.Context) error {
	beanId := c.Param("beanId")
	bean := services.GetBean(beanId)

	return c.JSON(http.StatusOK, bean)
}