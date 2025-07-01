package main

import (
	   "bytes"
	   "encoding/json"
	   "fmt"
	   "io"
	   "log"
	   "net/http"
	   "os"
	   "github.com/joho/godotenv"
)

// Request from user
type AskRequest struct {
	Question string `json:"question"`
}

// Gemini JSON request structure
type GeminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

// Gemini API response structure
type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func main() {
	   // Load environment variables from .env file
	   _ = godotenv.Load()
	   http.HandleFunc("/ask", askHandler)
	   fmt.Println("✅ Server running on http://localhost:8080")
	   log.Fatal(http.ListenAndServe(":8080", nil))
}

func askHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON input
	var userInput AskRequest
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil || userInput.Question == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Query Gemini with the user's question
	answer, err := queryGemini(userInput.Question)
	if err != nil {
		http.Error(w, "Gemini API error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"question": userInput.Question,
		"answer":   answer,
	})
}

func queryGemini(question string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY not set")
	}

	// Gemini API Endpoint
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s", apiKey)

	// Prepare request payload
	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": question},
				},
			},
		},
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse response
	respBody, _ := io.ReadAll(resp.Body)
	var geminiResp GeminiResponse
	err = json.Unmarshal(respBody, &geminiResp)
	if err != nil {
		return "", err
	}

	if len(geminiResp.Candidates) == 0 {
		return "", fmt.Errorf("no response from Gemini")
	}

	return geminiResp.Candidates[0].Content.Parts[0].Text, nil
}
