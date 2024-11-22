package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		logTime := start.Format("2006-01-02 15:04:05")

		logLine := fmt.Sprintf("%-19s | %-3d | %10v | %-15s | %-7s %q\n",
			logTime,
			wrapped.statusCode,
			time.Since(start),
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
		)

		logger := log.New(os.Stdout, "", 0)
		logger.Print(logLine)
	})
}
