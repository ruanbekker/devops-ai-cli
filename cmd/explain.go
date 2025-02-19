package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var explainCmd = &cobra.Command{
	Use:   "explain <query>",
	Short: "Ask OpenWebUI for an explanation",
	Long:  `Send a query to OpenWebUI and display the response in Markdown.`,
	Args:  cobra.ExactArgs(1), // Require exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]

		// Read API settings from config.yaml
		apiHost := viper.GetString("openwebui.host")
		apiKey := viper.GetString("openwebui.api_key")

		// If no API key in config, check environment variable
		if apiKey == "" {
			apiKey = os.Getenv("OPENWEB_API_KEY")
		}

		// Exit if no API key is found
		if apiHost == "" || apiKey == "" {
			fmt.Println("Error: OpenWebUI host and API key must be set in config.yaml or environment variable OPENWEB_API_KEY")
			os.Exit(1)
		}

		// Construct API request payload
		requestBody, err := json.Marshal(map[string]interface{}{
			"model": "gemma:2b",
			"messages": []map[string]string{
				{"role": "user", "content": query},
			},
		})
		if err != nil {
			fmt.Printf("Error creating JSON payload: %v\n", err)
			os.Exit(1)
		}

		// Send API request
		url := fmt.Sprintf("%s/api/chat/completions", apiHost)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			os.Exit(1)
		}

		// Set headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			os.Exit(1)
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
			fmt.Printf("Error parsing JSON response: %v\n", err)
			os.Exit(1)
		}

		if len(jsonResponse.Choices) == 0 {
			fmt.Println("No response received from OpenWebUI.")
			os.Exit(1)
		}

		// Extract the message content
		messageContent := jsonResponse.Choices[0].Message.Content

		// Render response using Glamour
		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(80),
		)
		if err != nil {
			fmt.Printf("Error initializing renderer: %v\n", err)
			os.Exit(1)
		}

		renderedOutput, err := renderer.Render(messageContent)
		if err != nil {
			fmt.Printf("Error rendering markdown: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(renderedOutput)
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}

