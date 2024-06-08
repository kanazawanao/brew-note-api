package services

import (
	"brew-note/src/database"
	"brew-note/src/models"
)

func GetProcessings() []models.Processing {
	var processings []models.Processing

	result := database.Handler.Find(&processings)

	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return processings
}
