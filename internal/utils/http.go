package utils

import (
	"encoding/json"
	"mable-to-facebook/internal/models"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, response models.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
