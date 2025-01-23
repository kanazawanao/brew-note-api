package services

import (
	"brew-note/src/database"
	"brew-note/src/models"
)

func CreateRecipeStep(recipeStep models.RecipeStep) models.RecipeStep {
	result := database.Handler.Create(&recipeStep)

	if err := result.Error; err != nil {
		panic(err.Error())
	}
	return recipeStep
}
