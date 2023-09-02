package services

import (
	"coffee-paws/src/database"
	"coffee-paws/src/models"
	"log"
)

func GetProfile(token string) models.User {
	claim, err := checkFirebaseJWT(token)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	var user = models.User{ID: claim.Id}
	result := database.Handler.FirstOrCreate(&user)
	if err := result.Error; err != nil {
		panic(err.Error())
	}
	// userにメールアドレスをつめる？
	

	return user
}
