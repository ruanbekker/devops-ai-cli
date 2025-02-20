package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/ruanbekker/devops-ai-cli/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

var optimizeFilePath string

var optimizeCmd = &cobra.Command{
	Use:   "optimize -f <file>",
	Short: "Optimize a code or configuration file using AI",
	Long: `Reads a code/configuration file (YAML, JSON, Python, Terraform, Shell, etc.)
and sends it to OpenWebUI API for optimization. The AI returns suggestions in Markdown format.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure file path is provided
		if optimizeFilePath == "" {
			fmt.Println("Error: Please specify a file with -f")
			os.Exit(1)
		}

		// Read API settings from config.yaml or environment variables
		apiHost := viper.GetString("openwebui.host")
		apiKey := viper.GetString("openwebui.api_key")
		aiModel := viper.GetString("openwebui.model")

		// If no API host in config, check environment variable
		if apiHost == "" {
			apiHost = os.Getenv("OPENWEB_API_HOST")
		}

		// If no API key in config, check environment variable
		if apiKey == "" {
			apiKey = os.Getenv("OPENWEB_API_KEY")
		}

		// Exit if no API key is found
		if apiHost == "" || apiKey == "" {
			fmt.Println("Error: OpenWebUI host and API key must be set in config.yaml or environment variable OPENWEB_API_KEY")
			os.Exit(1)
		}

		// Read the file content
		content, err := os.ReadFile(optimizeFilePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}

		// Get the file extension (to provide context to AI)
		fileType := detectFileType(optimizeFilePath)

		// Log debug information
		if viper.GetBool("debug") {
			logger.Log(fmt.Sprintf("optimize: using %s model for %s", aiModel, fileType))
		}

		// Send file content to OpenWebUI
		markdownResponse, err := sendToOpenWebUI(apiHost, apiKey, aiModel, string(content), fileType)
		if err != nil {
			fmt.Printf("Error from AI: %v\n", err)
			os.Exit(1)
		}

		// Render Markdown using Glamour
		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(80),
		)
		if err != nil {
			fmt.Printf("Error initializing renderer: %v\n", err)
			os.Exit(1)
		}

		renderedOutput, err := renderer.Render(markdownResponse)
		if err != nil {
			fmt.Printf("Error rendering markdown: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(renderedOutput)
	},
}

func init() {
	// Add flag for file input
	optimizeCmd.Flags().StringVarP(&optimizeFilePath, "file", "f", "", "Path to the file to optimize")
	rootCmd.AddCommand(optimizeCmd)
}

// detectFileType returns a file type based on the extension
func detectFileType(filePath string) string {
	switch {
	case hasExtension(filePath, ".yaml", ".yml"):
		return "Kubernetes YAML"
	case hasExtension(filePath, ".json"):
		return "JSON configuration"
	case hasExtension(filePath, ".tf"):
		return "Terraform script"
	case hasExtension(filePath, ".sh"):
		return "Shell script"
	case hasExtension(filePath, ".py"):
		return "Python script"
	default:
		return "Unknown format"
	}
}

// hasExtension checks if the file has one of the provided extensions
func hasExtension(filePath string, extensions ...string) bool {
	for _, ext := range extensions {
		if len(filePath) > len(ext) && filePath[len(filePath)-len(ext):] == ext {
			return true
		}
	}
	return false
}

// sendToOpenWebUI sends the file content to OpenWebUI API and returns Markdown suggestions
func sendToOpenWebUI(apiHost, apiKey, model, content, fileType string) (string, error) {
	// Construct API request payload
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "user", "content": fmt.Sprintf("Optimize this %s:\n\n%s", fileType, content)},
		},
	})
	if err != nil {
		return "", err
	}

	// Send API request
	url := fmt.Sprintf("%s/api/chat/completions", apiHost)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse JSON response
	var jsonResponse struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return "", err
	}

	if len(jsonResponse.Choices) == 0 {
		return "⚠️ No response received from OpenWebUI", nil
	}

	// Extract the AI-generated Markdown response
	return jsonResponse.Choices[0].Message.Content, nil
}

