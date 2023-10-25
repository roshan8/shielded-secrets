# ShieldedSecrets

## BE:

### API contracts:

1. GET /regions
   Description: Fetch all the available AWS region names
   Response: {"data": ["us-east1", ".."], "meta": ..}

2. GET /secrets?region=us-east1
   Description: Fetch all secret names from a region.
   Response: {"data": ["secret1", "secret2"], "meta": ..}

3. GET /secrets/{secret-name}?region=us-east1
   Description: Fetch all the key names of a secret
   Response: {"data": ["key1", "key2"], "meta": ..}

4. PUT /secrets/{secret-name}?region=us-east1
   Description: Add a new key to existing secret
   Payload: {"key": "value"} 
   Response: {"data": ["key1"], "meta": ..}

5. DELETE /secrets/{secret-name}?region=us-east1&key="keyName"
   Description: Delete a key
   Response: {"data": ["key1"], "meta": ..}

### Schema
TBD
