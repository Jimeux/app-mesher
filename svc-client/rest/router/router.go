package router

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/Jimeux/app-mesher/svc-client/rest/handlers"
)

type Handlers struct {
	Token   *handlers.TokenHandler
}

func Init(h *Handlers) chi.Router {
	r := chi.NewRouter()
	initMiddleware(r)
	initRoutes(r, h)
	return r
}

func initRoutes(r chi.Router, h *Handlers) {
	// ALB health check
	r.Get("/client", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// v1
	r.Route("/client/v1", func(r chi.Router) {
		r.Route("/tokens", func(r chi.Router) {
			r.Post("/", h.Token.GetToken)
		})
	})
}
