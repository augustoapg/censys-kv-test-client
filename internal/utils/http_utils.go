package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

func WriteResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := map[string]interface{}{
		"message": message,
	}

	encoder := json.NewEncoder(w)

	if os.Getenv("ENV") != "development" {
		encoder.SetIndent("", "  ")
	}
	encoder.Encode(response)
}
