package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	ServerPort         string `"server_port"`
	DBConnectionString string `"db_connection_string"`
	// Add more configuration fields as needed
}

var config Config

// LoadConfig loads the configuration from a file
func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Failed to unmarshal configuration: %v", err)
	}

	return config
}

// GetConfig returns the loaded configuration
func GetConfig() Config {
	return config
}
