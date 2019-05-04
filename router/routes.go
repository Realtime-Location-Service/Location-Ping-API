package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rls/ping-api/utils/error"
)

var router = chi.NewRouter()

type errResponse struct {
	Err *error.Err `json:"err"`
}
// Route returns the api router
func Route() http.Handler {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; utf8")
		w.WriteHeader(http.StatusNotFound)
		err := errResponse{error.NewErr(http.StatusNotFound, "Route not found!")}
		json.NewEncoder(w).Encode(err)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; utf8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errResponse{error.NewErr(http.StatusMethodNotAllowed, "Method not allowed!")}
		json.NewEncoder(w).Encode(err)
	})

	registerRoutes()

	return router
}

func registerRoutes() {
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/locations", locationHandler())
	})
}

// TODO: will be implemeneted
func locationHandler() http.Handler {
	return nil
}