package main

import (
	"log/slog"
	"os"

	"github.com/rmntim/go-url-shortener/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config
	config := config.MustLoad()

	// TODO: init logger
  logger := setupLogger(config.Env)

	// TODO: init db

	// TODO: init router

	// TODO: init server
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
  case envDev:
    logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
  case envProd:
    logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
