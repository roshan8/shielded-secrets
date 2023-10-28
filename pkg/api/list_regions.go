package api

import (
	"net/http"
	"shielded-secrets/pkg/respond"
	"shielded-secrets/vars"
)

func regionHandler(w http.ResponseWriter, r *http.Request) {
	respond.OK(w, vars.Regions, nil)
}
