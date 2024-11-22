package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unicode/utf8"

	_ "github.com/lib/pq"
)

func isTextFile(content []byte) bool {
	// Sprawdzenie, czy plik jest plikiem tekstowym (nie zawiera null byte)
	for _, b := range content {
		if b == 0 {
			// Jeśli znajdziemy null byte, oznacza to plik binarny
			return false
		}
	}
	return utf8.Valid(content) // Sprawdza, czy zawartość jest poprawnym tekstem UTF-8
}

func main() {
	// Tworzymy lub otwieramy plik logu
	logFile, err := os.OpenFile("process.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	defer logFile.Close()

	// Tworzymy logger, który zapisuje logi tylko do pliku
	logger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Parametry połączenia z bazą danych PostgreSQL
	connStr := "postgres://postgres:test@localhost:5432/my_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// Ścieżka do katalogu z plikami
	baseDir := "/media/pawel/storage/backup/"

	// Iteracja po katalogach
	err = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Printf("Error walking the path %v: %v\n", path, err)
			return err
		}

		// Sprawdzenie, czy jest to plik (pomijamy katalogi)
		if !info.IsDir() {
			// Odczytanie treści pliku
			content, err := os.ReadFile(path)
			if err != nil {
				logger.Printf("Error reading file %s: %v\n", path, err)
				return err
			}

			// Sprawdzenie, czy plik jest tekstowy
			if !isTextFile(content) {
				logger.Printf("Skipping binary file: %s\n", path)
				return nil
			}

			// Obliczanie hasha treści pliku
			hash := sha256.Sum256(content)

			// Przygotowanie zapytania do dodania pliku do bazy danych
			_, err = db.Exec(`
				INSERT INTO my_schema.files (file_name, content_hash, content)
				VALUES ($1, $2, $3)
				ON CONFLICT (content_hash) DO NOTHING
			`, path, fmt.Sprintf("%x", hash), string(content))
			if err != nil {
				logger.Printf("Error inserting file %s into database: %v\n", path, err)
				return err
			}

			logger.Printf("Successfully added file: %s\n", path)
		}
		return nil
	})

	if err != nil {
		logger.Fatal("Error walking the directory tree: ", err)
	}
}
