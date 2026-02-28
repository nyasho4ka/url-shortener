package delete

import (
	"log/slog"
	"net/http"
	resp "url-shortner/internal/lib/api/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type URLDeleter interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("invalid request"))
			return
		}

		err := urlDeleter.DeleteURL(alias)

		if err != nil {
			log.Error("failed to delete url", "err", err)
			render.JSON(w, r, resp.Error("failed to delete url"))
			return
		}

		log.Info("url deleted", "alias", alias)

		render.JSON(
			w, r, resp.OK(),
		)
	}
}
