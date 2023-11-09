package api

import (
	"shielded-secrets/vars"

	"github.com/go-chi/chi/v5"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Routes(router chi.Router) {
	// Endpoint for prometheus to scrape metrics
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
	router.Route("/regions", initRegionSubRoutes)
}

// Routes:
// /regions - list all regions
// /regions/{regionID} - list all secrets in a region
// /regions/{regionID}/{secretID} - get a secret within a region

func initRegionSubRoutes(router chi.Router) {
	router.Use(allowOnlyIPs(vars.AllowedIPs))
	router.Get("/", regionHandler)
	router.With(regionRequired).Route("/{regionID}", initRegionSecretSubRoutes)
}

func initRegionSecretSubRoutes(router chi.Router) {
	router.Get("/", listSecretsHandler)
	router.With(secretRequired).Route("/{secretID}", initSecretSubRoutes)
}

func initSecretSubRoutes(router chi.Router) {
	router.Get("/", getSecretHandler)
	router.Put("/", addSecretKeyHandler)
	router.Delete("/", deleteSecretKeyHandler)
}
