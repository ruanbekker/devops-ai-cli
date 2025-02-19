package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ruanbekker/devops-ai-cli/internal/logger"
)

var version = "dev"

var getVersion = func() string {
	configVersion := viper.GetString("version")
	if configVersion == "" {
		configVersion = version
	}
	return configVersion
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of devops-ai-cli",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Log("Running version command")
		if viper.GetBool("debug") {
			fmt.Println("DevOps AI CLI - Debug Mode")
		}
		fmt.Printf("DevOps AI CLI v%s\n", getVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

