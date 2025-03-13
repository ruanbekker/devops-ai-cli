package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/ruanbekker/devops-ai-cli/config"
)

var defaultConfig = `# DevOpsCLI Configuration
openwebui:
  host: "http://localhost:3000"
  api_key: ""
  model: "gemma:2b"

debug: false
`

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize DevOpsCLI with a default config file",
	Long:  `Creates a sample configuration file at ~/.config/devopscli/config.yaml.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetConfigPath()
		configDir := filepath.Dir(configPath)

		// Ensure directory exists
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			err := os.MkdirAll(configDir, 0755)
			if err != nil {
				fmt.Println("❌ Error creating config directory:", err)
				os.Exit(1)
			}
		}

		// Check if config already exists
		if _, err := os.Stat(configPath); err == nil {
			fmt.Println("⚠️  Config file already exists at", configPath)
			return
		}

		// Create default config file
		err := os.WriteFile(configPath, []byte(defaultConfig), 0644)
		if err != nil {
			fmt.Println("❌ Error writing config file:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Config initialized at", configPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

