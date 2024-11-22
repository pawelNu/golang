package server

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// HelloWorldHandler godoc
//
//	@Summary		Hello World Handler
//	@Description	Returns a hello world message
//	@Tags			test
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{string}	string	"Internal Server Error"
//	@Router			/hello-world [get]
func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Endpoint do wystawiania pliku - wyświetlanie pliku w przeglądarce
// @Summary Display a file
// @Description Endpoint do wyświetlania zawartości pliku w przeglądarce
// @Tags file
// @Produce text/plain
// @Success 200 {string} string "File content displayed"
// @Failure 404 {string} string "File not found"
// @Router /file/show [get]
func (s *Server) ShowFileHandler(w http.ResponseWriter, r *http.Request) {
	// Ścieżka do pliku na serwerze
	filePath := "./uploaded_files/your_file.txt"

	// Otwieramy plik do odczytu
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Ustawiamy odpowiedni typ content-type (w tym przypadku tekst)
	w.Header().Set("Content-Type", "text/plain")

	// Kopiujemy zawartość pliku do odpowiedzi HTTP
	io.Copy(w, file)
}

func (s *Server) DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Ścieżka do pliku na serwerze
	filePath := "./uploaded_files/your_file.txt"

	// Otwieramy plik do odczytu
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Ustawiamy nagłówek, aby przeglądarka traktowała odpowiedź jako plik do pobrania
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	w.Header().Set("Content-Type", "application/octet-stream")

	// Kopiujemy zawartość pliku do odpowiedzi HTTP
	io.Copy(w, file)
}
