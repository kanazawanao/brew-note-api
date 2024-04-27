package services

import (
	"brew-note/src/database"
	"brew-note/src/models"
)

func PostBean(store models.Bean) models.Bean {
	result := database.Handler.Create(&store)

	if err := result.Error; err != nil {
		panic(err.Error())
	}
	return store
}

func GetBeans(userId string) []models.Bean {
	var beans []models.Bean
	result := database.Handler.Where("user_id = ?", userId).Find(&beans)

	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return beans
}

func GetBean(id int) models.Bean {
	var bean = models.Bean{ID: id}
	result := database.Handler.First(&bean)
	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return bean
}
