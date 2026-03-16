package paperextractor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ajay-patidar-0/rag/internal/model"
)

func PaperImageToJson(encodedPaper string) (*model.ExamPaper, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-3-flash-preview:generateContent?key=%s",
		apiKey,
	)

	body := model.CreateRequestBody(encodedPaper)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, err
	}
	fmt.Printf("size of buffer is %d bytes\n", buf.Len())
	resp, err := http.Post(url, "application/json", &buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API Error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result model.GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("Empty response from model")
	}

	var resultData model.ExamPaper
	if err := json.Unmarshal([]byte(result.Candidates[0].Content.Parts[0].Text), &resultData); err != nil {
		return nil, err
	}

	fmt.Printf("course name is %v and exam year is %v length of the slice of questions %v", resultData.CourseName, resultData.ExamYear, len(resultData.Questions))

	return &resultData, nil
}
