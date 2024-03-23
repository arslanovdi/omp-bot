package logger

import (
	"log/slog"
	"os"
)

func InitializeLogger(level slog.Level) {
	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: level,
		}))

	slog.SetDefault(logger)
}
