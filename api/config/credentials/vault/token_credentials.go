package vault

import "github.com/dtcookie/gojson"

// TokenCredentials A credentials set of the `TOKEN` type.
type TokenCredentials struct {
	BaseCredentials `json:",type=TOKEN"`
	Token           string `json:"token"` // Token in the string format.
}

// UnmarshalJSON provides custom JSON deserialization
func (tc *TokenCredentials) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, tc)
}

// MarshalJSON provides custom JSON serialization
func (tc *TokenCredentials) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(tc)
}
