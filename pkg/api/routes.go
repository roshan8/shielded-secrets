package api

import (
	"net"
	"net/http"
	"shielded-secrets/vars"

	"github.com/go-chi/chi/v5"

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
	router.Get("/", listSecretsHandler)
	router.With(secretRequired).Route("/{secretID}", initSecretSubRoutes)
}

func initSecretSubRoutes(router chi.Router) {
	router.Get("/", getSecretHandler)
	router.Put("/", addSecretKeyHandler)
	router.Delete("/", deleteSecretKeyHandler)
}

func allowOnlyIPs(allowedIPs []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			allowed := false
			for _, ip := range allowedIPs {
				if clientIP == ip {
					allowed = true
					break
				}
			}

			if !allowed {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
