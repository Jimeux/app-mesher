package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func initMiddleware(r chi.Router) {
	r.Use(
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
		middleware.StripSlashes,
		middleware.Logger,
		middleware.Compress(5),
	)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set a maximum body size of 10MB
			r.Body = http.MaxBytesReader(w, r.Body, 1e+7)
			next.ServeHTTP(w, r)
		})
	})
}
