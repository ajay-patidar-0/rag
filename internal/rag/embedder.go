package rag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type EmbeddingResponse struct {
	Embedding struct {
		Values []float32 `json:"values"`
	} `json:"embedding"`
}

func GetEmbedding(text string) ([]float32, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-embedding-001:embedContent?key=%s",
		apiKey,
	)

	body := map[string]interface{}{
		"model": "models/gemini-embedding-001",
		"content": map[string]interface{}{
			"parts": []map[string]string{
				{"text": text},
			},
		},
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read the body to see the actual error message from Google
		var errRes interface{}
		json.NewDecoder(resp.Body).Decode(&errRes)
		return nil, fmt.Errorf("API error %s: %v", resp.Status, errRes)
	}

	var result EmbeddingResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Embedding.Values, nil
}
