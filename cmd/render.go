package cmd

import (
	"fmt"
	"os"
	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var filePath string

var renderCmd = &cobra.Command{
	Use: "render",
	Short: "Render a markdown file with Glow",
	Long: `Render a markdown file in the terminal using Glow.`,
	Run: func(cmd *cobra.Command, args []string) {
		if filePath == "" {
			fmt.Println("Error: Please specify a markdown file with -f")
			os.Exit(1)
		}

		// Read the markdown file
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}

		// Render markdown using Glamour
		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(80),
		)
		if err != nil {
			fmt.Printf("Error initializing renderer: %v\n", err)
			os.Exit(1)
		}

		rendered, err := renderer.Render(string(content))
		if err != nil {
			fmt.Printf("Error rendering markdown: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(rendered)

	},
}

func init(){
	renderCmd.Flags().StringVarP(&filePath, "file", "f", "", "Markdown file to render")
	rootCmd.AddCommand(renderCmd)
}
