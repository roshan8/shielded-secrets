package api

import (
	"context"
	"net/http"
	"shielded-secrets/pkg/respond"
	"shielded-secrets/vars"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Routes(router chi.Router) {
	// Endpoint for prometheus to scrape metrics
	router.Get("/metrics", promhttp.Handler().ServeHTTP)

	router.Get("/regions", regionHandler)
	router.With(regionRequired).Route("/{regionID}", initRegionSubRoutes)

	// router.Method(http.MethodGet, "/readyz", api.Handler(api2.ReadyzHandler))
	// router.Method(http.MethodGet, "/livez", api.Handler(api2.LivezHandler))
}

func initRegionSubRoutes(router chi.Router) {
	router.Get("/", fetchSecretsHandler)
	router.With(secretRequired).Route("/{secretID}", initSecretSubRoutes)
}

func initSecretSubRoutes(router chi.Router) {
	// router.Get("/", fetchSecretHandler)
	// router.Post("/", createSecretHandler)
	// router.Put("/", updateSecretHandler)
	// router.Delete("/", deleteSecretHandler)
}

func secretRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		regionID := chi.URLParam(r, vars.RegionID)
		ctx = context.WithValue(ctx, vars.RegionID, regionID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func regionRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		regionID := chi.URLParam(r, vars.RegionID)

		found := false
		for _, region := range vars.Regions {
			if region == regionID {
				found = true
				break
			}
		}
		if !found {
			respond.Fail(w, errors.Errorf("region %s not found", regionID))
			return
		}

		ctx = context.WithValue(ctx, vars.RegionID, regionID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
