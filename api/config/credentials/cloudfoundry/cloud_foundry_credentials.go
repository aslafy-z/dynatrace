package cloudfoundry

import (
	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/gojson"
)

// CloudFoundryCredentials Configuration for specific Cloud Foundry credentials.
type CloudFoundryCredentials struct {
	LoginURL           string                     `json:"loginUrl"`                     // The login URL of the Cloud Foundry foundation credentials.  The URL must be valid according to RFC 2396.  Leading or trailing whitespaces are not allowed.
	Metadata           *api.ConfigurationMetadata `json:"metadata,omitempty"`           // Metadata useful for debugging
	Password           *string                    `json:"password,omitempty"`           // The password of the Cloud Foundry foundation credentials.
	Active             *bool                      `json:"active,omitempty"`             // The monitoring is enabled (`true`) or disabled (`false`) for given credentials configuration.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	EndpointStatus     *EndpointStatus            `json:"endpointStatus,omitempty"`     // The status of the configured endpoint. ASSIGNED: The credentials are assigned to an ActiveGate which is responsible for processing. UNASSIGNED: The credentials are not yet assigned to an ActiveGate so there is currently no processing. DISABLED: The credentials have been disabled by the user. FASTCHECK_AUTH_ERROR: The credentials are invalid. FASTCHECK_TLS_ERROR: The endpoint TLS certificate is invalid. FASTCHECK_NO_RESPONSE: The endpoint did not return a result until the timeout was reached. FASTCHECK_INVALID_ENDPOINT: The endpoint URL was invalid. FASTCHECK_AUTH_LOCKED: The credentials seem to be locked. UNKNOWN: An unknown error occured.
	EndpointStatusInfo *string                    `json:"endpointStatusInfo,omitempty"` // The detailed status info of the configured endpoint.
	ID                 *string                    `json:"id,omitempty"`                 // The ID of the given credentials configuration.
	Name               string                     `json:"name"`                         // The name of the Cloud Foundry foundation credentials.  Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`. Leading or trailing whitespace is not allowed.
	Username           string                     `json:"username"`                     // The username of the Cloud Foundry foundation credentials.  Leading and trailing whitespaces are not allowed.
	ApiURL             string                     `json:"apiUrl"`                       // The URL of the Cloud Foundry foundation credentials.  The URL must be valid according to RFC 2396.  Leading or trailing whitespaces are not allowed.
}

// UnmarshalJSON provides custom JSON deserialization
func (cfc *CloudFoundryCredentials) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, cfc)
}

// MarshalJSON provides custom JSON serialization
func (cfc *CloudFoundryCredentials) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(cfc)
}
