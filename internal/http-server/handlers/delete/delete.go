package delete

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	resp "github.com/rmntim/go-url-shortener/internal/lib/api/response"
	"github.com/rmntim/go-url-shortener/internal/lib/logger/sl"
	"github.com/rmntim/go-url-shortener/internal/storage"
)

type URLDeleter interface {
	DeleteURL(alias string) error
}

func New(logger *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.delete.New"

		logger = logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			logger.Info("empty alias")
			render.JSON(w, r, resp.Error("invalid request"))
			return
		}

		if err := urlDeleter.DeleteURL(alias); err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				logger.Info("alias not found", slog.String("alias", alias))
				render.JSON(w, r, resp.Error("not found"))
				return
			}

			logger.Error("failed to delete url", sl.Err(err))
			render.JSON(w, r, resp.Error("internal server error"))
			return
		}

		logger.Info("url deleted", slog.String("alias", alias))

		render.JSON(w, r, resp.Ok())
	}
}
