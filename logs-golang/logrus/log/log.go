package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Globalna instancja loggera
var logger = createLogger()

// createLogger tworzy nowego loggera z rotacją plików
func createLogger() *logrus.Logger {
	log := logrus.New()

	// Ustawienie rotacji plików
	logFile := &lumberjack.Logger{
		Filename:   "logfile.log",
		MaxSize:    10, // MB
		MaxBackups: 5,  // Maksymalna liczba backupów
		MaxAge:     30, // Dni
		Compress:   false,
	}

	// Ustawienie wyjścia do logów na konsolę i do pliku
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	// Ustawienie niestandardowego formatu logowania
	log.SetFormatter(&CustomFormatter{})

	return log
}

// CustomFormatter definiuje niestandardowy format logowania
type CustomFormatter struct{}

// Format formatuje logi w żądanym formacie
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Kolory dla poziomów logowania
	var levelColor, fileLineColor string
	level := entry.Level.String()

	switch entry.Level {
	case logrus.InfoLevel:
		levelColor = "\033[34m" // Niebieski dla INFO
		level = "INFO"
	case logrus.WarnLevel:
		levelColor = "\033[33m" // Żółty dla WARN
		level = "WARN"
	case logrus.ErrorLevel:
		levelColor = "\033[31m" // Czerwony dla ERROR
		level = "ERROR"
	default:
		levelColor = "\033[0m" // Domyślny kolor
	}

	// Kolor dla pliku i numeru linii
	fileLineColor = "\033[36m" // Cyan

	// Formatowanie logu
	return []byte(
		fmt.Sprintf("%s %s[%s]%s %s[%s:%d]%s: %s\033[0m\n",
			entry.Time.Format("2006-01-02 15:04:05"),
			levelColor,         // Kolor dla poziomu logowania
			level,              // Poziom logowania
			"\033[0m",          // Reset koloru po poziomie
			fileLineColor,      // Kolor dla pliku i linii
			entry.Data["file"], // Nazwa pliku
			entry.Data["line"], // Numer linii
			"\033[0m",          // Reset koloru po pliku i linii
			entry.Message,      // Wiadomość logu
		)), nil
}

// logWithDetails loguje wiadomość z informacjami o pliku i numerze linii
func logWithDetails(level logrus.Level, message string) {
	_, file, line, ok := runtime.Caller(2) // Użycie 2, aby uzyskać wywołania z metody użytkownika
	if ok {
		// Używamy filepath.Base, aby uzyskać tylko nazwę pliku
		file = filepath.Base(file)
		logger.WithFields(logrus.Fields{
			"file": file,
			"line": line,
		}).Log(level, message)
	}
}

// Info loguje wiadomość z poziomem INFO
func Info(message string) {
	logWithDetails(logrus.InfoLevel, message)
}

// Warn loguje wiadomość z poziomem WARN
func Warn(message string) {
	logWithDetails(logrus.WarnLevel, message)
}

// Error loguje wiadomość z poziomem ERROR
func Error(message string) {
	logWithDetails(logrus.ErrorLevel, message)
}
