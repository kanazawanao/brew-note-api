package database

import (
	"time"
	"tripig/src/models"
	"tripig/test/e2e/data"
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
		placeType.Key = ""
		placeType.Name = ""
		now := time.Now()
		placeType.CreatedAt = &now
		DB.Create(&placeType)
		time.Sleep(1000 * time.Millisecond)
	}
}