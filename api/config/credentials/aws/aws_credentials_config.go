package aws

import (
	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/gojson"
)

// AWSCredentialsConfig Configuration of an AWS credentials.
type AWSCredentialsConfig struct {
	Label                       string                       `json:"label"`                                 // The name of the credentials.
	SupportingServicesToMonitor []AWSSupportingServiceConfig `json:"supportingServicesToMonitor,omitempty"` // A list of supporting services to be monitored.
	TaggedOnly                  bool                         `json:"taggedOnly"`                            // Monitor only resources which have specified AWS tags (`true`) or all resources (`false`).
	AuthenticationData          AWSAuthenticationData        `json:"authenticationData"`                    // A credentials for the AWS authentication.
	ID                          *string                      `json:"id,omitempty"`                          // The unique ID of the credentials.
	PartitionType               PartitionType                `json:"partitionType"`                         // The type of the AWS partition.
	TagsToMonitor               []AWSConfigTag               `json:"tagsToMonitor"`                         // A list of AWS tags to be monitored.  You can specify up to 10 tags.  Only applicable when the **taggedOnly** parameter is set to `true`.
	ConnectionStatus            *ConnectionStatus            `json:"connectionStatus,omitempty"`            // The status of the connection to the AWS environment.   * `CONNECTED`: There was a connection within last 10 minutes.  * `DISCONNECTED`: A problem occurred with establishing connection using these credentials. Check whether the data is correct.  * `UNINITIALIZED`: The successful connection has never been established for these credentials.
	Metadata                    *api.ConfigurationMetadata   `json:"metadata,omitempty"`                    // Metadata useful for debugging
}

// UnmarshalJSON provides custom JSON deserialization
func (acc *AWSCredentialsConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, acc)
}

// MarshalJSON provides custom JSON serialization
func (acc *AWSCredentialsConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(acc)
}
