package aws

import "github.com/dtcookie/gojson"

// AWSAuthenticationData A credentials for the AWS authentication.
type AWSAuthenticationData struct {
	KeyBasedAuthentication  *KeyBasedAuthentication  `json:"keyBasedAuthentication,omitempty"`  // The credentials for the key-based authentication.
	RoleBasedAuthentication *RoleBasedAuthentication `json:"roleBasedAuthentication,omitempty"` // The credentials for the role-based authentication.
	Type                    Type                     `json:"type"`                              // The type of the authentication: role-based or key-based.
}

// UnmarshalJSON provides custom JSON deserialization
func (aad *AWSAuthenticationData) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, aad)
}

// MarshalJSON provides custom JSON serialization
func (aad *AWSAuthenticationData) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(aad)
}
