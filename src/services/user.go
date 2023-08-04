package services

import (
	"tripig/src/database"
	"tripig/src/models"
)

type Users []models.User

func GetUsers() []models.User {
	var users Users
	result := database.Handler.Find(&users)

	if err := result.Error; err != nil {
		panic(err.Error())
	}

	return users
}

func PostUser(user models.User) {
	result := database.Handler.Create(&user)
	
	if err := result.Error; err != nil {
		panic(err.Error())
	}
}

func GetUser(id string) models.User {
	var user = models.User{ID: id}
	result := database.Handler.First(&user)
	if err := result.Error; err != nil {
		panic(err.Error())
	}
	
	return user
}