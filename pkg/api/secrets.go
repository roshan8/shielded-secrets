package api

import (
	"net/http"
	"shielded-secrets/pkg/respond"
	"shielded-secrets/vars"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

func fetchSecretsHandler(w http.ResponseWriter, r *http.Request) {
	var secretNames []string
	region := r.Context().Value(vars.RegionID).(string)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Error().Msgf("Error creating AWS session: %v", err)
		respond.Fail(w, err)
		return
	}

	client := secretsmanager.New(sess)

	// Define parameters for listing secrets.
	params := &secretsmanager.ListSecretsInput{}

	secrets, err := client.ListSecrets(params)
	if err != nil {
		log.Error().Msgf("Error listing secrets: %v", err)
		respond.Fail(w, err)
		return
	}

	for _, secret := range secrets.SecretList {
		secretNames = append(secretNames, *secret.Name)
	}

	respond.OK(w, secretNames, nil)
}
