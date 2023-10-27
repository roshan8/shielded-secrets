package api

import (
	"fmt"
	"net/http"
	"shielded-secrets/pkg/respond"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

func fetchSecretsHandler(w http.ResponseWriter, r *http.Request) {
	// region := chi.URLParam(r, "regionID")
	region := "us-east-1"
	var secretNames []string

	// Create a new AWS session.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Error().Msgf("Error creating AWS session: %v", err)
		respond.Fail(w, err)
		return
	}

	// Create a Secrets Manager client.
	client := secretsmanager.New(sess)

	// Define parameters for listing secrets.
	params := &secretsmanager.ListSecretsInput{}

	// List all the secrets.
	secrets, err := client.ListSecrets(params)
	if err != nil {
		log.Error().Msgf("Error listing secrets: %v", err)
		respond.Fail(w, err)
		return
	}

	// Loop through the secrets and print their names.
	for _, secret := range secrets.SecretList {
		fmt.Println(*secret.Name)
		secretNames = append(secretNames, *secret.Name)
	}

	respond.OK(w, secretNames, nil)
}
