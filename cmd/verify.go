package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// verifyToolsCmd checks if required tools are installed
var verifyToolsCmd = &cobra.Command{
	Use:   "verify tools",
	Short: "Check if required DevOps tools are installed",
	Long: `Reads the list of required tools from config.yaml and checks if they are installed 
on your local system. Displays a ✅ for installed tools and ❌ for missing ones.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read required tools from config.yaml
		requiredTools := viper.GetStringSlice("tools.required")
		if len(requiredTools) == 0 {
			fmt.Println("No tools defined in config.yaml")
			os.Exit(1)
		}

		fmt.Println("\n🔍 **Verifying Required Tools:**\n")

		// Check each tool and display result
		for _, tool := range requiredTools {
			if isToolInstalled(tool) {
				fmt.Printf("✅ %s\n", tool)
			} else {
				fmt.Printf("❌ %s (Not Installed)\n", tool)
			}
		}

		fmt.Println("")
	},
}

func init() {
	// ✅ Register under rootCmd instead of generateCmd
	rootCmd.AddCommand(verifyToolsCmd)
}

// isToolInstalled checks if a tool is available in PATH
func isToolInstalled(tool string) bool {
	_, err := exec.LookPath(tool)
	return err == nil
}

