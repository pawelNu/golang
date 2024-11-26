package log

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Globalna instancja loggera
var logger zerolog.Logger

// InitLogger inicjalizuje logger
func init() {
	// Ustawienia loggera
	zerolog.TimeFieldFormat = time.RFC3339

	// Konfiguracja rotacji logów
	fileLogger := &lumberjack.Logger{
		Filename:   "application.log",
		MaxSize:    1,     // MB
		MaxBackups: 3,     // Przechowuj maksymalnie 3 kopie zapasowe
		MaxAge:     30,    // Przechowuj logi przez 30 dni
		Compress:   false, // Kompresuj stare logi
	}

	multi := zerolog.MultiLevelWriter(os.Stdout, fileLogger)
	logger = zerolog.New(multi).With().Timestamp().Logger()
}

// logWithDetails loguje wiadomość z informacjami o pliku i numerze linii
func logWithDetails(level zerolog.Level, message string) {
	_, file, line, ok := runtime.Caller(2) // Użycie 2, aby uzyskać wywołania z metody użytkownika
	if ok {
		file = filepath.Base(file) // Uzyskanie tylko nazwy pliku
		switch level {
		case zerolog.DebugLevel:
			logger.Debug().Str("file", file).Int("line", line).Msg(message)
		case zerolog.InfoLevel:
			logger.Info().Str("file", file).Int("line", line).Msg(message)
		case zerolog.WarnLevel:
			logger.Warn().Str("file", file).Int("line", line).Msg(message)
		case zerolog.ErrorLevel:
			logger.Error().Str("file", file).Int("line", line).Msg(message)
		}
	}
}

// Info loguje wiadomość z poziomem INFO
func Info(message string) {
	logWithDetails(zerolog.InfoLevel, message)
}

// Warn loguje wiadomość z poziomem WARN
func Warn(message string) {
	logWithDetails(zerolog.WarnLevel, message)
}

// Error loguje wiadomość z poziomem ERROR
func Error(message string) {
	logWithDetails(zerolog.ErrorLevel, message)
}
