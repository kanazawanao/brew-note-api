package main

import (
	"brew-note/src/config"
	"brew-note/src/database"
	"brew-note/src/router"
)

func main() {
	// database conn
	database.NewSQLHandler()
	defer database.Handler.Close()
	router := router.NewRouter()
	// Start server
	router.Logger.Fatal(router.Start(":" + config.App.Port))
}
