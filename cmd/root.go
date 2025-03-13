package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ruanbekker/devops-ai-cli/config"
	"github.com/ruanbekker/devops-ai-cli/internal/logger"
)

var rootCmd = &cobra.Command{
	Use:   "devopscli",
	Short: "A DevOps CLI",
	Long:  `DevOps AI CLI is a simple, extensible Go CLI using Cobra and Viper.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("debug") {
			logger.Log("Debug mode is enabled")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run `devopscli help` for available commands")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug mode")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}
