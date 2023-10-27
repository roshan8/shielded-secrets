package api

import (
	"net/http"
	"shielded-secrets/pkg/respond"
)

func getSecretHandler(w http.ResponseWriter, r *http.Request) {
	secret, _ := getSecretAndAwsClient(r)

	sanitizedValues, er := stripSecretValues(secret)
	if er != nil {
		respond.Fail(w, er)
		return
	}

	respond.OK(w, sanitizedValues, nil)
}
