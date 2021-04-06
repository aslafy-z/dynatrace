package vault

import "github.com/dtcookie/gojson"

// UserPasswordCredentials A credentials set of the `USERNAME_PASSWORD` type.
type UserPasswordCredentials struct {
	BaseCredentials `json:",type=USERNAME_PASSWORD"`
	Password        string `json:"password"` // The password of the credential.
	User            string `json:"user"`     // The username of the credentials set.
}

// UnmarshalJSON provides custom JSON deserialization
func (upc *UserPasswordCredentials) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, upc)
}

// MarshalJSON provides custom JSON serialization
func (upc *UserPasswordCredentials) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(upc)
}
