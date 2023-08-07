package main

import (
	"coffee-paws/src/config"
	"coffee-paws/src/database"
	"coffee-paws/src/router"
)

func main() {
	// database conn
	database.NewSQLHandler()
	defer database.Handler.Close()
	router := router.NewRouter()
	// Start server
	router.Logger.Fatal(router.Start(":" + config.App.Port))
}
