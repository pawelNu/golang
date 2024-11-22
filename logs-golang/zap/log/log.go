package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"runtime"
)

// Globalna instancja loggera
var logger *zap.Logger

// InitLogger inicjalizuje logger
func init() {
	config := zap.NewProductionConfig()

	// Ustawienia kodera
	config.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   nil, // Nie używamy domyślnego kodera dla caller
	}

	// Ustawienia wyjścia do pliku
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	// Dodajemy writer do pliku
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config.EncoderConfig), // Wybór kodera JSON
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logFile)), // Wypisujemy zarówno na konsolę, jak i do pliku
		zapcore.InfoLevel, // Poziom logowania
	)

	logger = zap.New(core)
}

// logWithDetails loguje wiadomość z informacjami o pliku i numerze linii
func logWithDetails(level zapcore.Level, message string) {
	_, file, line, ok := runtime.Caller(2) // Użycie 2, aby uzyskać wywołania z metody użytkownika
	if ok {
		file = filepath.Base(file) // Uzyskanie tylko nazwy pliku
		switch level {
		case zapcore.DebugLevel:
			logger.Debug(message, zap.String("file", file), zap.Int("line", line))
		case zapcore.InfoLevel:
			logger.Info(message, zap.String("file", file), zap.Int("line", line))
		case zapcore.WarnLevel:
			logger.Warn(message, zap.String("file", file), zap.Int("line", line))
		case zapcore.ErrorLevel:
			logger.Error(message, zap.String("file", file), zap.Int("line", line))
		}
	}
}

// Info loguje wiadomość z poziomem INFO
func Info(message string) {
	logWithDetails(zapcore.InfoLevel, message)
}

// Warn loguje wiadomość z poziomem WARN
func Warn(message string) {
	logWithDetails(zapcore.WarnLevel, message)
}

// Error loguje wiadomość z poziomem ERROR
func Error(message string) {
	logWithDetails(zapcore.ErrorLevel, message)
}
