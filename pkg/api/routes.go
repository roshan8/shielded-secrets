package api

import (
	"shielded-secrets/vars"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Routes(router chi.Router) {
	// Endpoint for prometheus to scrape metrics
	router.Get("/", indexHandler)
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
	router.Get("/regions", regionHandler)
	router.With(regionRequired).Route("/{regionID}", initRegionSubRoutes)
}

func initRegionSubRoutes(router chi.Router) {
	router.Use(allowOnlyIPs(vars.AllowedIPs))

	if vars.BasicAuthUsername != "" && vars.BasicAuthPassword != "" {
		router.Use(middleware.BasicAuth("user", map[string]string{
			vars.BasicAuthUsername: vars.BasicAuthPassword,
		}))
	}

	router.Get("/", listSecretsHandler)
	router.With(secretRequired).Route("/{secretID}", initSecretSubRoutes)
}

func initSecretSubRoutes(router chi.Router) {
	router.Get("/", getSecretHandler)
	router.Put("/", addSecretKeyHandler)
	router.Delete("/", deleteSecretKeyHandler)
}
