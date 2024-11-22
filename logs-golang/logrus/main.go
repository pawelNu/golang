package main

import "logsLogrs/log"

func main() {
	log.Info("To jest informacja.")
	log.Warn("To jest ostrzeżenie.")
	log.Error("To jest błąd.")

	// Możesz dodać dodatkowe funkcjonalności w swoim programie
	log.Info("Aplikacja uruchomiona pomyślnie.")
}