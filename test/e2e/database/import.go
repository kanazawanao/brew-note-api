package database

import (
	"time"
	"tripig/src/models"
	"tripig/test/e2e/data"
)

func ImportData() {
	for _, id := range data.UserIDs {
		user := &models.User{}
		user.ID = id
		user.Name = ""
		now := time.Now()
		user.CreatedAt = now
		DB.Create(&user)
		time.Sleep(1000 * time.Millisecond)
	}
}