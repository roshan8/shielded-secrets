# ShieldedSecrets

## BE:

### API contracts:

1. GET /regions
   Description: Fetch all the available AWS region names
   Response: {"data": ["us-east1", ".."], "meta": ..}

2. GET /secrets/{regionID}/
   Description: Fetch all secret names from a region.
   Response: {"data": ["secret1", "secret2"], "meta": ..}

3. GET /secrets/{regionID}/{secret-name}/
   Description: Fetch all the key names of a secret
   Response: {"data": ["key1", "key2"], "meta": ..}

4. PUT /secrets/{regionID}/{secret-name}
   Description: Add a new key to existing secret
   Payload: {"key": "value"} 
   Response: {"data": ["key1"], "meta": ..}

5. DELETE /secrets/{regionID}/{secret-name}?key="keyName"
   Description: Delete a key
   Response: {"data": ["key1"], "meta": ..}

### TODO
- [ ] Fix the envelope for API response. (especially for meta)
- [ ] Support TLS with lets encrypt.
- [ ] Support basic auth.
- [ ] Support trusted IP.
- [ ] Add contributing.md
- [ ] Embed FE in Go and ship single binary
- [ ] Add Dockerfile and document instructions for how to deploy this app on prod.