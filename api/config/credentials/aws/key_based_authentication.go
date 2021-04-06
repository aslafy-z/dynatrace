package aws

import "github.com/dtcookie/gojson"

// KeyBasedAuthentication The credentials for the key-based authentication.
type KeyBasedAuthentication struct {
	AccessKey string `json:"accessKey"` // The ID of the access key.
	SecretKey string `json:"secretKey"` // The secret access key.
}

// UnmarshalJSON provides custom JSON deserialization
func (kba *KeyBasedAuthentication) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, kba)
}

// MarshalJSON provides custom JSON serialization
func (kba *KeyBasedAuthentication) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(kba)
}
