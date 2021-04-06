package azure

import (
	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/gojson"
)

// AzureCredentials Configuration of Azure app-level credentials.
type AzureCredentials struct {
	DirectoryID               string                     `json:"directoryId"`                  // The Directory ID (also referred to as Tenant ID)  The combination of Application ID and Directory ID must be unique.
	AutoTagging               bool                       `json:"autoTagging"`                  // The automatic capture of Azure tags is on (`true`) or off (`false`).
	Label                     string                     `json:"label"`                        // The unique name of the Azure credentials configuration.  Allowed characters are letters, numbers, and spaces. Also the special characters `.+-_` are allowed.
	Metadata                  *api.ConfigurationMetadata `json:"metadata,omitempty"`           // Metadata useful for debugging
	MonitorOnlyTagPairs       []CloudTag                 `json:"monitorOnlyTagPairs"`          // A list of Azure tags to be monitored.  You can specify up to 10 tags. A resource tagged with *any* of the specified tags is monitored.  Only applicable when the **monitorOnlyTaggedEntities** parameter is set to `true`.
	Active                    *bool                      `json:"active,omitempty"`             // The monitoring is enabled (`true`) or disabled (`false`).  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	AppID                     string                     `json:"appId"`                        // The Application ID (also referred to as Client ID)  The combination of Application ID and Directory ID must be unique.
	ID                        *string                    `json:"id,omitempty"`                 // The Dynatrace entity ID of the Azure credentials configuration.
	Key                       *string                    `json:"key,omitempty"`                // The secret key associated with the Application ID.  For security reasons, GET requests return this field as `null`.   Submit your key on creation or update of the configuration. If the field is omitted during an update, the old value remains unaffected.
	MonitorOnlyTaggedEntities bool                       `json:"monitorOnlyTaggedEntities"`    // Monitor only resources that have specified Azure tags (`true`) or all resources (`false`).
	SupportingServices        []AzureSupportingService   `json:"supportingServices,omitempty"` // A list of Azure supporting services to be monitored. For each service there's a sublist of its metrics and the metrics' dimensions that should be monitored. All of these elements (services, metrics, dimensions) must have corresponding static definitions on the server.
}

// UnmarshalJSON provides custom JSON deserialization
func (ac *AzureCredentials) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ac)
}

// MarshalJSON provides custom JSON serialization
func (ac *AzureCredentials) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ac)
}
