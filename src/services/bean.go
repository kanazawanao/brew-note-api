package services

import (
	"coffee-paws/src/database"
	"coffee-paws/src/models"
)

func PostBean(store models.Bean) models.Bean {
	result := database.Handler.Create(&store)

	if err := result.Error; err != nil {
		panic(err.Error())
	}
	return store
}

func GetBeans(storeId string) []models.Bean {
	var beans []models.Bean
	result := database.Handler.Where("store_id = ?", storeId).Find(&beans)

	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return beans
}


func GetBean(id string) models.Bean {
	var bean = models.Bean{ID: id}
	result := database.Handler.First(&bean)
	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return bean
}