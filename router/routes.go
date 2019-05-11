package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rls/ping-api/pkg/config"
	"github.com/rls/ping-api/pkg/location"
	"github.com/rls/ping-api/store/repo"
	"github.com/rls/ping-api/svc/cache"
	"github.com/rls/ping-api/utils/errors"
)

var router = chi.NewRouter()

type errResponse struct {
	Err *errors.Err `json:"err"`
}

// Route returns the api router
func Route() http.Handler {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; utf8")
		w.WriteHeader(http.StatusNotFound)
		err := errResponse{errors.NewErr(http.StatusNotFound, "Route not found!")}
		json.NewEncoder(w).Encode(err)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; utf8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errResponse{errors.NewErr(http.StatusMethodNotAllowed, "Method not allowed!")}
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

func locationHandler() http.Handler {
	var locationSvc location.Service
	cacheSvc := cache.NewCacheService(config.AppCfg().CacheType)
	locationSvc = location.NewService(repo.NewLocation(cacheSvc))
	locationSvc = location.NewValidationMiddleware(locationSvc)

	return location.MakeHandler(locationSvc)
}
