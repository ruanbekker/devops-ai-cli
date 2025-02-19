package cmd

import (
	"testing"
)

func TestRootCmd(t *testing.T) {
	if rootCmd.Use != "devopscli" {
		t.Errorf("expected root command to be 'devopscli', got %s", rootCmd.Use)
	}
}

