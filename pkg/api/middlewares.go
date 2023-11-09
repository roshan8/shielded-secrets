package api

import (
	"context"
	"net"
	"net/http"

	"shielded-secrets/pkg/respond"
	"shielded-secrets/vars"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

func regionRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		regionID := chi.URLParam(r, vars.RegionID)

		if !isValidRegion(regionID) {
			respond.Fail(w, errors.Errorf("region %s not found", regionID))
			return
		}

		ctx = context.WithValue(ctx, vars.RegionID, regionID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func secretRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		region := r.Context().Value(vars.RegionID).(string)
		secretID := chi.URLParam(r, vars.SecretID)

		client, err := getAWSClient(region)
		if err != nil {
			respond.Fail(w, err)
			return
		}

		secretData, err := getSecretData(client, secretID)
		if err != nil {
			respond.Fail(w, err)
			return
		}

		ctx = context.WithValue(ctx, vars.SecretID, secretID)
		ctx = context.WithValue(ctx, vars.Secret, secretData)
		ctx = context.WithValue(ctx, vars.AwsClient, client)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func allowOnlyIPs(allowedIPs []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			// allow all if no ips are specified
			if len(allowedIPs) == 0 {
				next.ServeHTTP(w, r)
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
