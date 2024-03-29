package main

import (
	"log/slog"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	cfg "github.com/rmntim/go-url-shortener/internal/config"
	del "github.com/rmntim/go-url-shortener/internal/http-server/handlers/delete"
	"github.com/rmntim/go-url-shortener/internal/http-server/handlers/redirect"
	"github.com/rmntim/go-url-shortener/internal/http-server/handlers/url/save"
	loggerMw "github.com/rmntim/go-url-shortener/internal/http-server/middleware/logger"
	"github.com/rmntim/go-url-shortener/internal/lib/logger/sl"
	"github.com/rmntim/go-url-shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := cfg.MustLoad()

	logger := setupLogger(config.Env)

	logger.Info("Starting url-shortener", slog.String("env", config.Env))
	logger.Debug("Debug messages are enabled")

	storage, err := sqlite.New(config.StoragePath)
	if err != nil {
		logger.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(loggerMw.New(logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// FIXME: not using proper auth middleware, should be JWT or smth
	router.Route("/api/url", func(r chi.Router) {
		r.Use(middleware.BasicAuth("url-shortener", map[string]string{
			config.User: config.Password,
		}))

		r.Post("/", save.New(logger, storage))
		r.Delete("/{alias}", del.New(logger, storage))
	})

	router.Get("/{alias}", redirect.New(logger, storage))

	logger.Info("starting server", slog.String("address", config.Address))
	srv := &http.Server{
		Addr:         config.Address,
		Handler:      router,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
		IdleTimeout:  config.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error("failed to start server")
	}

	logger.Error("server stopped")
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
