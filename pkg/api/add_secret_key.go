package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"shielded-secrets/pkg/respond"

	"github.com/rs/zerolog/log"
)

type secretKey struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func addSecretKeyHandler(w http.ResponseWriter, r *http.Request) {
	region, secretID := getRegionAndSecretID(r)

	var sk secretKey
	if err := json.NewDecoder(r.Body).Decode(&sk); err != nil {
		respond.Fail(w, err)
		return
	}

	if sk.Key == "" {
		respond.Fail(w, errors.New("key can not be nil"))
		return
	}

	if sk.Value == "" {
		respond.Fail(w, errors.New("value can not be nil"))
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

	if _, ok := secretData[sk.Key]; ok {
		log.Info().Msgf("Key '%s' already exists in the secret.\n", sk.Key)
		respond.Fail(w, errors.New("key already exists in the secret"))
		return
	}

	secretData[sk.Key] = sk.Value

	if err := updateSecretData(client, secretID, secretData); err != nil {
		log.Warn().Msgf("Error updating secret:", err)
		respond.Fail(w, errors.New("key not found in the secret"))
		return
	}

	msg := fmt.Sprintf("Key '%s' has been added to the secret '%s'", sk.Key, secretID)
	respond.OK(w, msg, nil)

	log.Info().Msgf("Key '%s' has been added to the secret '%s'.\n", sk.Key, secretID)
}
