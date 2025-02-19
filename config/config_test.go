package config

import (
	"testing"

	"github.com/spf13/viper"
)

// TestInitConfig ensures InitConfig() runs without error
func TestInitConfig(t *testing.T) {
	viper.Reset() // Reset Viper to avoid config interference

	// Call InitConfig()
	InitConfig()

	// Check that the default value for "debug" is correctly set
	if viper.GetBool("debug") != false {
		t.Errorf("Expected debug to be false, got %v", viper.GetBool("debug"))
	}
}

