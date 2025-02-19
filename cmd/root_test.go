package cmd

import (
	"testing"
)

func TestRootCmd(t *testing.T) {
	if rootCmd.Use != "starter" {
		t.Errorf("expected root command to be 'starter', got %s", rootCmd.Use)
	}
}

