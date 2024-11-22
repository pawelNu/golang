package utils

import (
	"encoding/json"
	"net/http"
	"webserwis/utils/log"
)

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
    // Ustawienie nagłówka odpowiedzi
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    // Utworzenie mapy dla odpowiedzi JSON
    errorResponse := map[string]string{
        "error": message,
    }

    // Kodowanie odpowiedzi do JSON i wysłanie jej
    if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
        log.Error("Nie udało się zakodować odpowiedzi JSON: %v", err)
    }
}

