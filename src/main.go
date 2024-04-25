package main

import (
	"brew-note/src/config"
	"brew-note/src/database"
	"brew-note/src/router"
	"fmt"
)

func main() {
	// database conn
	database.NewSQLHandler()
	defer database.Handler.Close()
	router := router.NewRouter()
	fmt.Print("test")
	// Start server
	router.Logger.Fatal(router.Start(":" + config.App.Port))
}
