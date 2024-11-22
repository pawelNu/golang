package utils

import (
	"encoding/json"
	"net/http"
	"webserwis/utils/log"
)

func ErrorResponse(w http.ResponseWriter, r *http.Request) {
	msg := "Invalid request"
	log.Error(msg)
	// http.Error(w, msg, http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	errorResponse := map[string]string{
		"error": msg,
	}
	json.NewEncoder(w).Encode(errorResponse)
}
