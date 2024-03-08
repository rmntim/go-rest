package main

import (
	"log/slog"
	"os"

	"github.com/rmntim/go-url-shortener/internal/config"
	"github.com/rmntim/go-url-shortener/internal/lib/logger/sl"
	"github.com/rmntim/go-url-shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := config.MustLoad()

  logger := setupLogger(config.Env)

  logger.Info("Starting url-shortener", slog.String("env", config.Env))
  logger.Debug("Debug messages are enabled")

  storage, err := sqlite.New(config.StoragePath)
  if err != nil {
    logger.Error("failed to init storage", sl.Err(err))
    os.Exit(1)
  }

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
