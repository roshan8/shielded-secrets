package api

import (
	"context"
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
