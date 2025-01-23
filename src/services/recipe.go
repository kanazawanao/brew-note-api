package services

import (
	"brew-note/src/database"
	"brew-note/src/models"
)

func CreateRecipe(recipe models.Recipe) models.Recipe {
	result := database.Handler.Create(&recipe)

	if err := result.Error; err != nil {
		panic(err.Error())
	}
	return recipe
}

func GetRecipes(userId string) []models.Recipe {
	var recipes []models.Recipe
	result := database.Handler.Where("user_id = ?", userId).Find(&recipes)

	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return recipes
}

func GetRecipe(id int) models.Recipe {
	var recipe = models.Recipe{ID: id}
	result := database.Handler.First(&recipe)
	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return recipe
}
