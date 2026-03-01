package utils

import (
	"encoding/json"
	"net/http"
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
