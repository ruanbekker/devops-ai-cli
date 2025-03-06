package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/charmbracelet/glamour"
	"github.com/ruanbekker/devops-ai-cli/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

// Path to store conversation history
var sessionFile = filepath.Join(os.TempDir(), "devopscli_session.json")

// Store conversation ID (simulated via local storage)
var conversationID string

var queryCmd = &cobra.Command{
	Use:   "query <message>",
	Short: "Ask OpenWebUI a question and maintain conversation context",
	Long: `Send a question to OpenWebUI and get a response.
Use --cid "<conversation-id>" to continue a previous conversation.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]

		// Read API settings from config.yaml or environment variables
		apiHost := viper.GetString("openwebui.host")
		apiKey := viper.GetString("openwebui.api_key")
		aiModel := viper.GetString("openwebui.model")

		if apiHost == "" {
			apiHost = os.Getenv("OPENWEB_API_HOST")
		}
		if apiKey == "" {
			apiKey = os.Getenv("OPENWEB_API_KEY")
		}

		if apiHost == "" || apiKey == "" {
			fmt.Println("Error: OpenWebUI host and API key must be set in config.yaml or environment variables.")
			os.Exit(1)
		}

		// Read existing conversation history if --cid is provided
		history := []map[string]string{}
		if conversationID != "" {
			history = loadConversationHistory(conversationID)
		}

		// Append the new user query to the history
		history = append(history, map[string]string{"role": "user", "content": message})

		// Debug logging
		if viper.GetBool("debug") {
			logger.Log(fmt.Sprintf("query: using %s model, conversation ID: %s", aiModel, conversationID))
		}

		// Send query to OpenWebUI
		response, err := sendQueryToOpenWebUI(apiHost, apiKey, aiModel, history)
		if err != nil {
			fmt.Printf("Error from OpenWebUI: %v\n", err)
			os.Exit(1)
		}

		// Append AI response to history
		history = append(history, map[string]string{"role": "assistant", "content": response})

		// Save the updated history with a new "conversation ID"
		newCID := saveConversationHistory(history)

		// Render Markdown response
		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(80),
		)
		if err != nil {
			fmt.Printf("Error initializing renderer: %v\n", err)
			os.Exit(1)
		}

		renderedOutput, err := renderer.Render(response)
		if err != nil {
			fmt.Printf("Error rendering markdown: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(renderedOutput)

		// Show the new conversation ID
		fmt.Printf("\n🆔 **Conversation ID**: %s\n", newCID)
	},
}

func init() {
	queryCmd.Flags().StringVarP(&conversationID, "cid", "c", "", "Continue a conversation with a conversation ID")
	rootCmd.AddCommand(queryCmd)
}

// sendQueryToOpenWebUI sends a conversation to OpenWebUI API and returns the AI response
func sendQueryToOpenWebUI(apiHost, apiKey, model string, history []map[string]string) (string, error) {
	// Prepare API request payload
	payload := map[string]interface{}{
		"model":    model,
		"messages": history,
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Send request
	url := fmt.Sprintf("%s/api/chat/completions", apiHost)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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

	// Extract response content
	if len(jsonResponse.Choices) == 0 {
		return "⚠️ No response received from OpenWebUI", nil
	}
	return jsonResponse.Choices[0].Message.Content, nil
}

// saveConversationHistory saves a conversation and returns a new CID
func saveConversationHistory(history []map[string]string) string {
	// Generate a fake "conversation ID" (incremental number)
	cid := fmt.Sprintf("%d", len(history))

	// Save the history to a session file
	file, err := os.Create(sessionFile)
	if err != nil {
		fmt.Println("Error saving conversation:", err)
		return cid
	}
	defer file.Close()

	jsonData, err := json.Marshal(history)
	if err != nil {
		fmt.Println("Error encoding conversation:", err)
		return cid
	}
	file.Write(jsonData)

	return cid
}

// loadConversationHistory loads a conversation by CID
func loadConversationHistory(cid string) []map[string]string {
	file, err := os.Open(sessionFile)
	if err != nil {
		// No history found, return empty
		return []map[string]string{}
	}
	defer file.Close()

	var history []map[string]string
	json.NewDecoder(file).Decode(&history)
	return history
}

