package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func RequestLogger(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

			defer func() {
				if err := recover(); err != nil {
					log.Error("panic recovered",
						"method", r.Method,
						"path", r.URL.Path,
						"error", err,
					)
					http.Error(rec, "internal server error", http.StatusInternalServerError)
				}

				log.Info("http_request",
					"method", r.Method,
					"path", r.URL.Path,
					"status", rec.status,
					"duration_ms", time.Since(start).Milliseconds(),
					"remote_addr", r.RemoteAddr,
					"user_agent", r.UserAgent(),
				)
			}()

			next.ServeHTTP(rec, r)
		})
	}
}
