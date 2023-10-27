package api

// stripSecretValues takes a JSON string containing secret key-value pairs
// and returns an array of keys with the values stripped out.
func stripSecretValues(secretMap map[string]interface{}) ([]string, error) {

	// Extract keys into an array.
	var keys []string
	for k := range secretMap {
		keys = append(keys, k)
	}

	return keys, nil
}
