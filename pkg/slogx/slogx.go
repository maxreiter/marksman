package slogx

import (
	"log/slog"
	"os"
)

func New(debug bool) *slog.Logger {
	var handler slog.Handler
	options := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}

	switch debug {
	case true:
		options.Level = slog.LevelDebug
		handler = slog.NewTextHandler(os.Stdout, options)
	case false:
		handler = slog.NewJSONHandler(os.Stdout, options)
	}

	return slog.New(handler)
}
