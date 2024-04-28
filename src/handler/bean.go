package handler

import (
	"brew-note/src/models"
	"brew-note/src/services"
	"log"
	"net/http"
	"strconv"

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

	token := c.Request().Header.Get("Authorization")
	claim, err := services.CheckFirebaseJWT(token)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	bean := models.Bean{
		UserId:         claim.UserId,
		ProductionArea: s.ProductionArea,
		Kind:           s.Kind,
		RoastLevelId:   s.RoastLevelId,
		Price:          0,
		Gram:           0,
	}

	res := services.PostBean(bean)
	return c.JSON(http.StatusOK, res)
}

// e.Get("/beans", GetBeans)
func GetBeans(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	claim, err := services.CheckFirebaseJWT(token)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	beans := services.GetBeans(claim.UserId)

	return c.JSON(http.StatusOK, beans)
}

// e.Get("/beans/:beanId", GetBean)
func GetBean(c echo.Context) error {
	beanId := c.Param("beanId")
	id, _ := strconv.Atoi(beanId)
	bean := services.GetBean(id)

	return c.JSON(http.StatusOK, bean)
}
