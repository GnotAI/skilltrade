package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// LoggerMiddleware logs incoming requests and their response status using zap.
func LoggerMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logFields := []zap.Field{
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("remote_addr", r.RemoteAddr),
			}

			// Capture response status
			wr := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(wr, r)

			logFields = append(logFields,
				zap.Int("status", wr.statusCode),
				zap.Duration("duration", time.Since(start)),
			)

			logger.Info("HTTP request", logFields...)
		})
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
