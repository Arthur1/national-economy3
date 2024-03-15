package logger

import (
	"log/slog"
	"os"

	"github.com/phsym/console-slog"
)

type Config struct {
	UseConsoleHandler bool
}

func New(cfg *Config) *slog.Logger {
	var handler slog.Handler
	if cfg.UseConsoleHandler {
		handler = console.NewHandler(os.Stderr, nil)
	} else {
		handler = slog.NewJSONHandler(os.Stderr, nil)
	}
	logger := slog.New(handler)
	return logger
}
