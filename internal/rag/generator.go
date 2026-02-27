package rag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
		FinishReason string `json:"finishReason"`
		Index        int    `json:"index"`
	} `json:"candidates"`
	UsageMetadata struct {
		PromptTokenCount     int `json:"promptTokenCount"`
		CandidatesTokenCount int `json:"candidatesTokenCount"`
		TotalTokenCount      int `json:"totalTokenCount"`
	} `json:"usageMetadata"`
}

func AskLLm(query string, context string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=%s",
		apiKey,
	)

	body := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{
						"text": "Answer based only on this context: " + context + "\n\nQuestion: " + query,
					},
				},
			},
		},
		"generationConfig": map[string]interface{}{
			"temperature": 0.1, // Low temperature is better for RAG accuracy
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", &buf)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	var result GeminiResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Candidates[0].Content.Parts[0].Text, nil

}
