package ui

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"webserwis/utils/log"
)

// UiHandler zwraca funkcję obsługującą statyczne pliki i routing
func UiHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ścieżka do pliku
		uiDir := "./frontend/dist"
		filePath := filepath.Join(uiDir, r.URL.Path)

		log.Info("Requesting file:", filePath)

		// Sprawdź, czy plik istnieje
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Error("File does not exist:", filePath)
			// Jeśli plik nie istnieje, zwróć index.html
			http.ServeFile(w, r, filepath.Join(uiDir, "index.html"))
			return
		}

		contentType := getMimeType(filePath)
		// Ustaw nagłówek Content-Type na odpowiedni typ MIME
		w.Header().Set("Content-Type", contentType)
		log.Info("Serving file:", filePath, "with Content-Type:", contentType)

		// Jeśli plik istnieje, zwróć go
		http.ServeFile(w, r, filePath)
	})
}

// Funkcja do ustalania odpowiedniego typu MIME na podstawie rozszerzenia pliku
func getMimeType(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	default:
		return "application/octet-stream"
	}
}
