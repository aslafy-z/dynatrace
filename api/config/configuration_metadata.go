package api

// ConfigurationMetadata is Metadata useful for debugging
type ConfigurationMetadata struct {
	CurrentConfigurationVersions []string `json:"currentConfigurationVersions,omitempty"` // A Sorted list of string version numbers of the configuration
	ClusterVersion               string   `json:"clusterVersion,omitempty"`               // Dynatrace server version
	ConfigurationVersions        []int64  `json:"configurationVersions,omitempty"`        // A Sorted list of the version numbers of the configuration
}

// String provides a string representation in JSON format for debugging purposes
func (cm ConfigurationMetadata) String() string {
	return MarshalJSON(&cm)
}
