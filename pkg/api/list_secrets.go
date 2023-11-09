package api

import (
	"net/http"

	"shielded-secrets/pkg/respond"
	"shielded-secrets/vars"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

func listSecretsHandler(w http.ResponseWriter, r *http.Request) {
	region := r.Context().Value(vars.RegionID).(string)

	client, err := getAWSClient(region)
	if err != nil {
		respond.Fail(w, err)
		return
	}

	secrets, err := client.ListSecrets(&secretsmanager.ListSecretsInput{})
	if err != nil {
		log.Error().Msgf("Error listing secrets: %v", err)
		respond.Fail(w, err)
		return
	}

	var secretNames []string
	for _, secret := range secrets.SecretList {
		secretNames = append(secretNames, *secret.Name)
	}

	respond.OK(w, secretNames, nil)
}
