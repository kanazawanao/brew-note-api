package database

import (
	"brew-note/src/models"
	"brew-note/test/e2e/data"
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
