package middleware

// TODO napisać od nowa
// TODO i przetestować od nowa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"webserwis/utils"
)

// TODO do poprawy do zrobienia jwt token
// Prosta struktura na dane logowania
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Tokeny w pamięci (prosta implementacja; w rzeczywistości użyj bazy danych lub JWT)
var validTokens = make(map[string]string) // token -> username
var validUsers = map[string]string{       // username -> password
	"user1": "password123",
	"user2": "securepass",
}

// Funkcja generująca token (w rzeczywistości użyj JWT)
func generateToken(username string) string {
	return fmt.Sprintf("%s:%d", username, time.Now().UnixNano())
}

// LoginHandler godoc
//
//	@Summary		Authorization endpoint
//	@Description	Send authorization data to receive token
//	@Tags			authorization
//	@Accept			json
//	@Produce		json
//	@Param			requestBody	body		Credentials	true	"User data payload"
//	@Success		200			{object}	map[string]string
//	@Failure		400			{string}	string	"Failed to parse JSON"
//	@Failure		405			{string}	string	"Invalid request method"
//	@Router			/login/ [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Sprawdzenie danych logowania
	if password, ok := validUsers[creds.Username]; !ok || password != creds.Password {
		utils.ErrorResponse(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generowanie tokena
	token := generateToken(creds.Username)
	validTokens[token] = creds.Username

	// Zwrócenie tokena w odpowiedzi
	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Middleware do autoryzacji
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Pobranie tokena (zakładamy schemat "Bearer <token>")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		username, ok := validTokens[token]
		if !ok {
			utils.ErrorResponse(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Wstawienie username do kontekstu żądania (opcjonalne)
		r.Header.Set("X-User", username)

		// Kontynuuj obsługę żądania
		next.ServeHTTP(w, r)
	})
}
