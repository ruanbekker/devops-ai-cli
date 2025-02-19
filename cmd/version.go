package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ruanbekker/go-cli-starter/internal/logger"
)

// Default version
var version = "dev"

// Function to fetch the version (can be mocked in tests)
var getVersion = func() string {
	configVersion := viper.GetString("version")
	if configVersion == "" {
		configVersion = version
	}
	return configVersion
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of go-cli-starter",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Log("Running version command")
		// Get the version from Viper (config.yaml), fallback to default if empty
		//configVersion := viper.GetString("version")
		//if configVersion == "" {
		//	configVersion = version
		//}
		if viper.GetBool("debug") {
			fmt.Println("CLI Starter - Debug Mode")
		}
		// fmt.Printf("CLI Starter v%s\n", configVersion)
		fmt.Printf("CLI Starter v%s\n", getVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

