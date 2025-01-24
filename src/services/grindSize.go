package services

import (
	"brew-note/src/database"
	"brew-note/src/models"
)

func GetGrindSizes() []models.GrindSize {
	var grindSizes []models.GrindSize

	result := database.Handler.Find(&grindSizes)

	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return grindSizes
}
