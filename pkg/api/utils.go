package api

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

// stripSecretValues takes a JSON string containing secret key-value pairs
// and returns an array of keys with the values stripped out.
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
