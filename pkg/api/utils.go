package api

import (
	"net/http"

	"shielded-secrets/vars"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

// stripSecretValues takes a JSON string containing secret key-value pairs
// and returns an array of keys with the values stripped out.
func stripSecretValues(secretMap map[string]interface{}) ([]string, error) {
	var keys []string

	for k := range secretMap {
		keys = append(keys, k)
	}

	return keys, nil
}

func isValidRegion(regionID string) bool {
	for _, region := range vars.Regions {
		if region == regionID {
			return true
		}
	}
	return false
}

func getRegionAndSecretID(r *http.Request) (string, string) {
	region := r.Context().Value(vars.RegionID).(string)
	secretID := r.Context().Value(vars.SecretID).(string)
	return region, secretID
}

func getSecretAndAwsClient(r *http.Request) (map[string]interface{}, *secretsmanager.SecretsManager) {
	secret := r.Context().Value(vars.Secret).(map[string]interface{})
	awsClient := r.Context().Value(vars.AwsClient).(*secretsmanager.SecretsManager)
	return secret, awsClient
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
