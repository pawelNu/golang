package middleware

// TODO napisać od nowa
// TODO i przetestować od nowa

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
	"webserwis/utils"

	jwt "github.com/golang-jwt/jwt/v5"
)

// TODO do poprawy do zrobienia jwt token
// Prosta struktura na dane logowania
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type LoginResponse struct {
	Username   string `json:"user_name" example:"user name"`
	ExpireDate string `json:"expire_date" example:"2024-11-27 01:31:43"`
	Token      string `json:"token" example:"aaa.bbb.ccc"`
}

var validUsers = map[string]string{
	"user1": "password123",
	"user2": "securepass",
}

var jwtSecret = []byte("your-very-secret-key")

// LoginHandler godoc
//
//	@Summary		Authorization endpoint
//	@Description	Send authorization data to receive token
//	@Tags			authorization
//	@Accept			json
//	@Produce		json
//	@Param			requestBody	body		Credentials	true	"User data payload"
//	@Success		200			{object}	LoginResponse
//	@Failure		400			{string}	string	"Failed to parse JSON"
//	@Failure		405			{string}	string	"Invalid request method"
//	@Router			/login [post]
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
	token, expirationTime, err := generateJWT(creds.Username, 9*time.Hour)
	if err != nil {
		utils.ErrorResponse(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Zwrócenie tokena, użytkownika i daty wygaśnięcia
	// response := map[string]interface{}{
	// 	"user":        creds.Username,
	// 	"expire_date": expirationTime.Format(time.DateTime), // Format ISO 8601
	// 	"token":       token,
	// }
	response := LoginResponse{
		Username:   creds.Username,
		ExpireDate: expirationTime.Format(time.DateTime),
		Token:      token,
	}
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
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Weryfikacja tokena
		claims, err := validateJWT(tokenString)
		if err != nil {
			utils.ErrorResponse(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Wstawienie username do nagłówka (opcjonalne)
		r.Header.Set("X-User", claims.Username)

		// Kontynuacja obsługi żądania
		next.ServeHTTP(w, r)
	})
}

// generateJWT generuje token JWT dla użytkownika oraz zwraca czas wygaśnięcia
func generateJWT(username string, duration time.Duration) (string, time.Time, error) {
	expirationTime := time.Now().Add(duration)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", time.Time{}, err
	}

	return signedToken, expirationTime, nil
}

// validateJWT sprawdza poprawność tokena i zwraca dane z niego
func validateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Upewniamy się, że używana jest właściwa metoda podpisywania
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
