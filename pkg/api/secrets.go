package api

import (
	"encoding/json"
	"net/http"
	"shielded-secrets/pkg/respond"
	"shielded-secrets/vars"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

func listSecretsHandler(w http.ResponseWriter, r *http.Request) {
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

func getSecretHandler(w http.ResponseWriter, r *http.Request) {
	region := r.Context().Value(vars.RegionID).(string)
	secretID := r.Context().Value(vars.SecretID).(string)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Error().Msgf("Error creating AWS session: %v", err)
		respond.Fail(w, err)
		return
	}

	client := secretsmanager.New(sess)

	// Define parameters for getting secret value.
	params := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}

	secretValue, err := client.GetSecretValue(params)
	if err != nil {
		log.Error().Msgf("Error getting secret value: %v", err)
		respond.Fail(w, err)
		return
	}

	sanitizedValues, er := stripSecretValues(*secretValue.SecretString)
	if er != nil {
		log.Error().Msgf("Error stripping secret values: %v", err)
		respond.Fail(w, err)
		return
	}

	respond.OK(w, sanitizedValues, nil)
}

func stripSecretValues(secretValue string) ([]string, error) {

	// Unmarshal the JSON string into a map.
	var secretMap map[string]interface{}
	if err := json.Unmarshal([]byte(secretValue), &secretMap); err != nil {
		log.Error().Msgf("Error unmarshalling secret value: %v", err)
		return nil, err
	}

	// Extract keys into an array.
	var keys []string
	for k := range secretMap {
		keys = append(keys, k)
	}

	return keys, nil
}
