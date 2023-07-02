package main

import (
	"tripig/src/app/config"
	"tripig/src/app/database"
	"tripig/src/app/router"
)

func main() {
	// database conn
	database.NewSQLHandler()
	defer database.Handler.Close()
	router := router.NewRouter()
	// Start server
	router.Logger.Fatal(router.Start(":" + config.Config.Port))
}
