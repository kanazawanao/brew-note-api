package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"tripig/src/app/utils"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port    string
	LogFile string
	Db      DbConfig
}

type DbConfig struct {
	User         string
	Password     string
	Host         string
	Port         string
	Name         string
	MaxIdleConns int
	MaxOpenConns int
}

var Config AppConfig

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	maxIdleConnsStr := os.Getenv("DB_MAX_IDLE_CONNS")
	maxIdleConns, _ := strconv.Atoi(maxIdleConnsStr)
	maxOpenConnsStr := os.Getenv("DB_MAX_OPEN_CONNS")
	maxOpenConns, _ := strconv.Atoi(maxOpenConnsStr)

	Config = AppConfig{
		Port:    os.Getenv("PORT"),
		LogFile: os.Getenv("LOG_FILE"),
		Db: DbConfig{
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			Name:         os.Getenv("DB_NAME"),
			MaxIdleConns: maxIdleConns,
			MaxOpenConns: maxOpenConns,
		},
	}

	fmt.Print("*****************", Config.Port)
}
