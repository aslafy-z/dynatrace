package vault

import "github.com/dtcookie/gojson"

// Credentials A set of credentials for synthetic monitors.
type Credentials interface {
	Initialize(*BaseCredentials) // Initialize duplicates the properties of the other AbstractConditionKey into this instance. It also serves as a unique method for structs implementing the interface this abstract class is derived from
	Implementors() []Credentials // Implementors provides automatic JSON unmarshalling with a list of struct implementing (directly or indirectly) ConditionKey
}

// BaseCredentials A set of credentials for synthetic monitors.
type BaseCredentials struct {
	Description     string           `json:"description"`               // A short description of the credentials set..
	ID              *string          `json:"id,omitempty"`              // The ID of the credentials set.
	Name            string           `json:"name"`                      // The name of the credentials set.
	OwnerAccessOnly *bool            `json:"ownerAccessOnly,omitempty"` // The credentials set is available to every user (`false`) or to owner only (`true`).
	Type            *CredentialsType `json:"type,omitempty"`            // Defines the actual set of fields depending on the value. See one of the following objects:  * `CERTIFICATE` -> CertificateCredentials  * `USERNAME_PASSWORD` -> UserPasswordCredentials  * `TOKEN` -> TokenCredentials
}

// Initialize duplicates the properties of the other BaseCredentials into this instance
// It also serves as a unique method for structs implementing the interface this abstract class is derived from
func (bc *BaseCredentials) Initialize(other *BaseCredentials) {
	bc.Description = other.Description
	bc.ID = other.ID
	bc.Name = other.Name
	bc.OwnerAccessOnly = other.OwnerAccessOnly
	bc.Type = other.Type
}

// Implementors provides automatic JSON unmarshalling with a list of struct implementing (directly or indirectly) Credentials
func (bc *BaseCredentials) Implementors() []Credentials {
	return []Credentials{
		new(CertificateCredentials),
		new(TokenCredentials),
		new(UserPasswordCredentials),
	}
}

// UnmarshalJSON provides custom JSON deserialization
func (bc *BaseCredentials) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, bc)
}

// MarshalJSON provides custom JSON serialization
func (bc *BaseCredentials) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(bc)
}
