package logger

import (
	"bytes"
	"log"
	"testing"

	"github.com/spf13/viper"
)

// Capture log output for testing
func captureLogOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil) // Reset log output

	f()
	return buf.String()
}

// TestLog ensures Log() prints debug messages when debug mode is enabled
func TestLog(t *testing.T) {
	viper.Reset()
	viper.Set("debug", true)

	output := captureLogOutput(func() {
		Log("Test message")
	})

	if output == "" {
		t.Errorf("Expected log output, but got empty string")
	}
}

