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

func getSecretHandler(w http.ResponseWriter, r *http.Request) {
	region, secretID := getRegionAndSecretID(r)

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

	sanitizedValues, er := stripSecretValues(secretData)
	if er != nil {
		respond.Fail(w, err)
		return
	}

	respond.OK(w, sanitizedValues, nil)
}

func getRegionAndSecretID(r *http.Request) (string, string) {
	region := r.Context().Value(vars.RegionID).(string)
	secretID := r.Context().Value(vars.SecretID).(string)
	return region, secretID
}

func getAWSClient(region string) (*secretsmanager.SecretsManager, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Error().Msgf("Error creating AWS session: %v", err)
		return nil, err
	}

	return secretsmanager.New(sess), nil
}

func getSecretValue(client *secretsmanager.SecretsManager, secretID string) (*secretsmanager.GetSecretValueOutput, error) {
	params := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}

	return client.GetSecretValue(params)
}
