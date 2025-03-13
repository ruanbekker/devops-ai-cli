package config

import (
	"fmt"
  "os"
  "path/filepath"

	"github.com/spf13/viper"
	"github.com/ruanbekker/devops-ai-cli/internal/logger"
)

// GetConfigPath returns the config path, defaulting to ~/.config/devopscli/config.yaml
func GetConfigPath() string {
	defaultConfigPath := filepath.Join(os.Getenv("HOME"), ".config", "devopscli", "config.yaml")

	// Allow override with environment variable
	if envConfigPath := os.Getenv("DEVOPSCLI_CONFIG_LOCATION"); envConfigPath != "" {
		return envConfigPath
	}
	return defaultConfigPath
}

// InitConfig initializes configuration
func InitConfig() {
  configPath := GetConfigPath()

  // Ensure directory exists
  configDir := filepath.Dir(configPath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}

  viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")   // File format

  // Default values
  viper.SetDefault("openwebui.host", "http://localhost:3000")
	viper.SetDefault("openwebui.api_key", "")
  viper.SetDefault("openwebui.model", "gemma:2b")
  viper.SetDefault("debug", false)

  // Read config file if available
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using default config, no config file found.")
	}

	if viper.GetBool("debug") {
		logger.Log("Configuration initialized with debug mode enabled")
    logger.Log(fmt.Sprintf("Configuration loaded from: %s", configPath))
	}
}

