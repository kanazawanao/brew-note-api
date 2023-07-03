package config

import (
	"os"
	"strconv"
	"tripig/src/utils"
)

type AppConfig struct {
	Port    string
	LogFile string
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
var DB DbConfig

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	maxIdleConnsStr := os.Getenv("DB_MAX_IDLE_CONNS")
	maxIdleConns, _ := strconv.Atoi(maxIdleConnsStr)
	maxOpenConnsStr := os.Getenv("DB_MAX_OPEN_CONNS")
	maxOpenConns, _ := strconv.Atoi(maxOpenConnsStr)

	Config = AppConfig{
		Port:    os.Getenv("PORT"),
		LogFile: os.Getenv("LOG_FILE"),
	}

	DB = DbConfig{
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		Name:         os.Getenv("DB_NAME"),
		MaxIdleConns: maxIdleConns,
		MaxOpenConns: maxOpenConns,
	}
}
