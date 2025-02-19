package logger

import (
	"log"
	"github.com/spf13/viper"
)

// Log prints messages only when debug mode is enabled
func Log(message string) {
	if viper.GetBool("debug") {
		log.Println("[DEBUG]", message)
	}
}

