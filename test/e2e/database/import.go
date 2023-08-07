package database

import (
	"coffee-paws/src/models"
	"coffee-paws/test/e2e/data"
	"time"
)

func ImportUserData() {
	for _, id := range data.UserIDs {
		user := &models.User{}
		user.ID = id
		user.Name = ""
		now := time.Now()
		user.CreatedAt = &now
		DB.Create(&user)
		time.Sleep(1000 * time.Millisecond)
	}
}

func ImportPlaceTypeDate() {
	for _, id := range data.PlaceTypeIds {
		placeType := &models.PlaceType{}
		placeType.ID = id
		placeType.Key = id + "key"
		placeType.Name = id + "name"
		now := time.Now()
		placeType.CreatedAt = &now
		DB.Create(&placeType)
		time.Sleep(1000 * time.Millisecond)
	}
}