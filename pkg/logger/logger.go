package logger

import (
	"log/slog"
	"os"
)

func NewSlogHandler() slog.Logger {
	log := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(log)

	return *logger
}
