package services

import (
	"coffee-paws/src/database"
	"coffee-paws/src/models"
	"coffee-paws/src/utils"
)

func PostStore(store models.Store) models.Store {
	id := utils.GenerateULID()
	store.ID = id;
	result := database.Handler.Create(&store)

	if err := result.Error; err != nil {
		panic(err.Error())
	}
	return store
}