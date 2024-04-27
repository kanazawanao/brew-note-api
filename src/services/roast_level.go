package services

import (
	"brew-note/src/database"
	"brew-note/src/models"
)

func GetRoastLevels() []models.RoastLevel {
	var roastLevels []models.RoastLevel
	result := database.Handler.Find(&roastLevels)

	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return roastLevels
}