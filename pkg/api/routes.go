package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Routes(router chi.Router) {
	// Endpoint for prometheus to scrape metrics
	router.Get("/", indexHandler)
	router.Get("/metrics", promhttp.Handler().ServeHTTP)

	router.Get("/regions", regionHandler)
	router.With(regionRequired).Route("/{regionID}", initRegionSubRoutes)

	// router.Method(http.MethodGet, "/readyz", api.Handler(api2.ReadyzHandler))
	// router.Method(http.MethodGet, "/livez", api.Handler(api2.LivezHandler))
}

func initRegionSubRoutes(router chi.Router) {
	router.Get("/", listSecretsHandler)
	router.With(secretRequired).Route("/{secretID}", initSecretSubRoutes)
}

func initSecretSubRoutes(router chi.Router) {
	router.Get("/", getSecretHandler)
	router.Put("/", addSecretKeyHandler)
	router.Delete("/", deleteSecretKeyHandler)
}
