package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"runtime/debug"
)

var cfg *Config

func LoadConfig() {
	loadDotEnvFile()
	readConfigFile()
	loadEnvVariables()

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to unmarshal config into struct: %s\n", err)
	}
}

func GetConfig() *Config {
	if cfg == nil {
		LoadConfig()
	}
	return cfg
}

func loadDotEnvFile() {
	currentPath := getCallerPath()
	if err := godotenv.Load(currentPath + "/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
}

func getCallerPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatalf(
			"Failed to find current path: %v (%v)\n",
			err,
			string(debug.Stack()),
		)
	}
	return currentPath + "/../.."
}

func readConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(getCallerPath())

	if err := viper.ReadInConfig(); err != nil {
		if errors.Is(err, viper.ConfigFileNotFoundError{}) {
			log.Println("config file not found, relying on the environment variables")
		} else {
			log.Fatalf("Fatal error reading config file: %v\n", err)
		}
	}
}

func loadEnvVariables() {
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()
	_ = viper.BindEnv("app.name", "APP_NAME")
	_ = viper.BindEnv("app.env", "APP_ENV")
	_ = viper.BindEnv("app.host", "APP_HOST")
	_ = viper.BindEnv("app.port", "APP_PORT")
	_ = viper.BindEnv("log.show_source", "LOG_SHOW_SOURCE")
}
