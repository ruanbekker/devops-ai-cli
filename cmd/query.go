package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/charmbracelet/glamour"
	"github.com/ruanbekker/devops-ai-cli/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

// Path to store conversation history
var sessionFile = filepath.Join(os.Getenv("HOME"), ".devopscli_sessions.json")

// Flags
var conversationID string
var listConversations bool

// Define a structure for conversations
type Conversation struct {
	ID      int                    `json:"id"`
	History []map[string]string    `json:"history"`
	Query   string                 `json:"query"`
}

// Define a slice to store conversations
type Conversations struct {
	List []Conversation `json:"conversations"`
}

var queryCmd = &cobra.Command{
	Use:   "query <message>",
	Short: "Ask OpenWebUI a question and maintain conversation context",
	Long: `Send a question to OpenWebUI and get a response.
Use --cid "<conversation-id>" to continue a previous conversation.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// If --list is used, show all conversations
		if listConversations {
			listStoredConversations()
			return
		}

		if len(args) == 0 {
			fmt.Println("Error: Please provide a query or use --list to view previous conversations.")
			os.Exit(1)
		}

		message := args[0]

		// Read API settings
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

		// Load conversation history if --cid is used
		history := []map[string]string{}
		conversationNumber := 0
		if conversationID != "" {
			history, conversationNumber = loadConversationByID(conversationID)
		}

		// Append new user query
		history = append(history, map[string]string{"role": "user", "content": message})

		// Debug log
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

		// Save updated conversation history
		newCID := saveConversation(history, conversationNumber, message)

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
		fmt.Printf("\n🆔 **Conversation ID**: %d\n", newCID)
	},
}

func init() {
	queryCmd.Flags().StringVarP(&conversationID, "cid", "c", "", "Continue a conversation with a conversation ID")
	queryCmd.Flags().BoolVarP(&listConversations, "list", "l", false, "List previous conversations")
	rootCmd.AddCommand(queryCmd)
}

// sendQueryToOpenWebUI sends a conversation to OpenWebUI API and returns the AI response
func sendQueryToOpenWebUI(apiHost, apiKey, model string, history []map[string]string) (string, error) {
	payload := map[string]interface{}{
		"model":    model,
		"messages": history,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/api/chat/completions", apiHost)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

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
	return jsonResponse.Choices[0].Message.Content, nil
}

// saveConversation saves a conversation and returns its ID
func saveConversation(history []map[string]string, existingCID int, query string) int {
	conversations := loadAllConversations()
	conversationID := existingCID

	if existingCID == 0 {
		conversationID = len(conversations.List) + 1
		conversations.List = append(conversations.List, Conversation{ID: conversationID, History: history, Query: query})
	} else {
		for i, conv := range conversations.List {
			if conv.ID == existingCID {
				conversations.List[i].History = history
			}
		}
	}

	jsonData, err := json.MarshalIndent(conversations, "", "  ")
	if err != nil {
		fmt.Println("Error encoding conversation:", err)
		return conversationID
	}

	err = os.WriteFile(sessionFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error saving conversation:", err)
	}

	return conversationID
}

// loadConversationByID retrieves a specific conversation
func loadConversationByID(cid string) ([]map[string]string, int) {
	conversations := loadAllConversations()
	id, err := strconv.Atoi(cid)
	if err != nil {
		return nil, 0
	}

	for _, conv := range conversations.List {
		if conv.ID == id {
			return conv.History, conv.ID
		}
	}
	return nil, 0
}

// loadAllConversations retrieves all stored conversations
func loadAllConversations() Conversations {
	data, err := os.ReadFile(sessionFile)
	if err != nil {
		return Conversations{}
	}

	var conversations Conversations
	json.Unmarshal(data, &conversations)
	return conversations
}

// listStoredConversations lists all stored conversations
func listStoredConversations() {
	conversations := loadAllConversations()
	if len(conversations.List) == 0 {
		fmt.Println("No previous conversations found.")
		return
	}

	fmt.Println("\n📝 **Previous Conversations:**")
	for _, conv := range conversations.List {
		fmt.Printf("🆔 %d: %s\n", conv.ID, conv.Query)
	}
}

