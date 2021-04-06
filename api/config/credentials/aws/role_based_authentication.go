package aws

import "github.com/dtcookie/gojson"

// RoleBasedAuthentication The credentials for the role-based authentication.
type RoleBasedAuthentication struct {
	AccountID  string  `json:"accountId"`            // The ID of the Amazon account.
	ExternalID *string `json:"externalId,omitempty"` // The external ID token for setting an IAM role.   You can obtain it with the `GET /aws/iamExternalId` request.
	IamRole    string  `json:"iamRole"`              // The IAM role to be used by Dynatrace to get monitoring data.
}

// UnmarshalJSON provides custom JSON deserialization
func (rba *RoleBasedAuthentication) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, rba)
}

// MarshalJSON provides custom JSON serialization
func (rba *RoleBasedAuthentication) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(rba)
}
