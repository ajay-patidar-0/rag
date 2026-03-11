package utils

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func RespondJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func RespondError(w http.ResponseWriter, statusCode int, message string) {
	RespondJson(w, statusCode, message)
}

func RespondSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	RespondJson(w, statusCode, data)
}

func ImageToBase64(imageUrl string) (string, error) {
	resp, err := http.Get(imageUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(imageBytes), nil
}

// just temporary use of image
func ImageTo64() (string, error) {
	imageBytes, err := os.ReadFile("./output_pages/page_1.jpg")
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(imageBytes), nil
}
