package redirect

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	resp "github.com/rmntim/go-url-shortener/internal/lib/api/response"
	"github.com/rmntim/go-url-shortener/internal/lib/logger/sl"
	"github.com/rmntim/go-url-shortener/internal/storage"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=URLGetter
type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(logger *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.redirect.New"

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

		url, err := urlGetter.GetURL(alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				logger.Info("alias not found", slog.String("alias", alias))
				render.JSON(w, r, resp.Error("not found"))
				return
			}

			logger.Error("failed to get url", sl.Err(err))
			render.JSON(w, r, resp.Error("internal server error"))
			return
		}

		logger.Info("url found", slog.String("url", url))

		http.Redirect(w, r, url, http.StatusFound)
	}
}
