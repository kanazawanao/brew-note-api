package services

import (
	"coffee-paws/src/database"
	"coffee-paws/src/models"
)

func PostStore(store models.Store) models.Store {
	result := database.Handler.FirstOrCreate(&store)

	if err := result.Error; err != nil {
		panic(err.Error())
	}
	return store
}


type Stores []models.Store

func GetStores() []models.Store {
	var stores Stores
	result := database.Handler.Find(&stores)

	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return stores
}