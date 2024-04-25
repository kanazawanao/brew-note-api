package config

import (
	"brew-note/src/utils"
	"os"
	"strconv"
)

type AppConfig struct {
	Port                 string
	LogFile              string
	GoogleMapApiKey      string
	OpenWeatherMapApiKey string
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

var App AppConfig
var DB DbConfig

func init() {
	LoadConfig()
	utils.LoggingSettings(App.LogFile)
}

func LoadConfig() {
	maxIdleConnsStr := os.Getenv("DB_MAX_IDLE_CONNS")
	maxIdleConns, _ := strconv.Atoi(maxIdleConnsStr)
	maxOpenConnsStr := os.Getenv("DB_MAX_OPEN_CONNS")
	maxOpenConns, _ := strconv.Atoi(maxOpenConnsStr)

	App = AppConfig{
		Port:                 os.Getenv("PORT"),
		LogFile:              os.Getenv("LOG_FILE"),
		GoogleMapApiKey:      os.Getenv("GOOGLE_MAP_API_KEY"),
		OpenWeatherMapApiKey: os.Getenv("OPEN_WEATHER_MAP_API_KEY"),
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
