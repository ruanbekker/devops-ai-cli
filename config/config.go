package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/ruanbekker/go-cli-starter/internal/logger"
)

// InitConfig initializes configuration
func InitConfig() {
	viper.SetConfigName("config") // Config filename
	viper.SetConfigType("yaml")   // File format
	viper.AddConfigPath(".")      // Search in current directory

	viper.SetDefault("debug", false)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using default config, no config file found.")
	}

	if viper.GetBool("debug") {
		logger.Log("Configuration initialized with debug mode enabled")
	}
}

