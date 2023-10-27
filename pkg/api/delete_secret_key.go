package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"shielded-secrets/pkg/respond"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

func deleteSecretKeyHandler(w http.ResponseWriter, r *http.Request) {
	region, secretID := getRegionAndSecretID(r)

	keyToDelete := r.URL.Query().Get("key")
	if keyToDelete == "" {
		respond.Fail(w, errors.New("key can not be nil"))
		return
	}

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

	if _, ok := secretData[keyToDelete]; !ok {
		log.Info().Msgf("Key '%s' not found in the secret.\n", keyToDelete)
		respond.Fail(w, errors.New("key not found in the secret"))
		return
	}

	delete(secretData, keyToDelete)

	if err := updateSecretData(client, secretID, secretData); err != nil {
		log.Warn().Msgf("Error updating secret:", err)
		respond.Fail(w, errors.New("key not found in the secret"))
		return
	}

	msg := fmt.Sprintf("Key '%s' has been deleted from the secret '%s'.\n", keyToDelete, secretID)
	respond.OK(w, msg, nil)

	log.Info().Msgf("Key '%s' has been deleted from the secret '%s'.\n", keyToDelete, secretID)
}

func getSecretData(client *secretsmanager.SecretsManager, secretID string) (map[string]interface{}, error) {
	secret, err := getSecretValue(client, secretID)
	if err != nil {
		return nil, err
	}

	secretData := make(map[string]interface{})
	if err := json.Unmarshal([]byte(*secret.SecretString), &secretData); err != nil {
		log.Warn().Msgf("Error parsing secret value:", err)
		return nil, err
	}

	return secretData, nil
}

func updateSecretData(client *secretsmanager.SecretsManager, secretID string, secretData map[string]interface{}) error {
	updatedSecretValue, err := json.Marshal(secretData)
	if err != nil {
		log.Warn().Msgf("Error marshalling updated secret data:", err)
		return err
	}

	updateInput := &secretsmanager.UpdateSecretInput{
		SecretId:     aws.String(secretID),
		SecretString: aws.String(string(updatedSecretValue)),
	}

	if _, err := client.UpdateSecret(updateInput); err != nil {
		return err
	}

	return nil
}
