package main

import (
	"tripig/src/config"
	"tripig/src/database"
	"tripig/src/router"
)

func main() {
	// database conn
	database.NewSQLHandler()
	defer database.Handler.Close()
	router := router.NewRouter()
	// Start server
	router.Logger.Fatal(router.Start(":" + config.Config.Port))
}
