package logger

import (
	"log/slog"
	"os"
	"strings"
)

func New() *slog.Logger {
	level := new(slog.LevelVar)

	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		level.Set(slog.LevelDebug)
	case "warn":
		level.Set(slog.LevelWarn)
	case "error":
		level.Set(slog.LevelError)
	default:
		level.Set(slog.LevelInfo)
	}
	env := strings.ToLower(os.Getenv("ENV"))
	var handler slog.Handler

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     level,
			AddSource: true,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     level,
			AddSource: false,
		})
	}

	return slog.New(handler)
}
