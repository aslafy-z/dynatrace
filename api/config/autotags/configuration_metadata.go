package autotags

import "github.com/dtcookie/gojson"

// ConfigurationMetadata Metadata useful for debugging
type ConfigurationMetadata struct {
	CurrentConfigurationVersions []string `json:"currentConfigurationVersions,omitempty"` // A Sorted list of string version numbers of the configuration.
	ClusterVersion               *string  `json:"clusterVersion,omitempty"`               // Dynatrace server version.
	ConfigurationVersions        []int64  `json:"configurationVersions,omitempty"`        // A Sorted list of the version numbers of the configuration.
}

// UnmarshalJSON provides custom JSON deserialization
func (cm *ConfigurationMetadata) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, cm)
}

// MarshalJSON provides custom JSON serialization
func (cm *ConfigurationMetadata) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(cm)
}
