package response

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, message string, status int) {

	responseBody := ResponseError{Status: http.StatusText(status), Message: message}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(responseBody)
}
