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

// e.POST("/recipes", PostRecipe)
func PostRecipe(c echo.Context) error {
	s := new(openapi.CreateRecipe)
	if err := c.Bind(s); err != nil {
		log.Printf("err %v", err.Error())
		return c.String(http.StatusBadRequest, "bad request")
	}

	token := c.Request().Header.Get("Authorization")
	claim, err := services.CheckFirebaseJWT(token)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return c.String(http.StatusUnauthorized, "unauthorized")
	}

	recipe := models.Recipe{
		UserId:              claim.UserId,
		Title:               s.Title,
		Description:         s.Description,
		GrindSizeId:         int(s.GrindSizeId),
		ExtractionEquipment: s.ExtractionEquipment,
		CoffeeType:          s.CoffeeType,
		WaterTemperature:    int(s.WaterTemperature),
		BeanDose:            int(s.BeanDose),
	}
	res := services.CreateRecipe(recipe)
	for _, step := range s.Steps {
		recipeStep := models.RecipeStep{
			RecipeId:    res.ID,
			StepNumber:  int(step.StepNumber),
			Seconds:     int(step.Seconds),
			AmountWater: int(step.AmountWater),
			Memo:        step.Memo,
		}
		services.CreateRecipeStep(recipeStep)
	}
	return c.JSON(http.StatusOK, res)
}

// e.Get("/recipes", GetRecipes)
func GetRecipes(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	claim, err := services.CheckFirebaseJWT(token)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	recipes := services.GetRecipes(claim.UserId)
	return c.JSON(http.StatusOK, recipes)
}

// e.Get("/recipes/:id", GetRecipe)
func GetRecipe(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("err %v", err.Error())
		return c.String(http.StatusBadRequest, "bad request")
	}
	recipe := services.GetRecipe(id)
	return c.JSON(http.StatusOK, recipe)
}
