package aws

import "github.com/dtcookie/gojson"

// AWSSupportingServiceConfig A supporting service to be monitored.
type AWSSupportingServiceConfig struct {
	MonitoredMetrics []AWSSupportingServiceMetric `json:"monitoredMetrics"` // A list of metrics to be monitored for this service.
	Name             string                       `json:"name"`             // The name of the supporting service.
}

// UnmarshalJSON provides custom JSON deserialization
func (assc *AWSSupportingServiceConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, assc)
}

// MarshalJSON provides custom JSON serialization
func (assc *AWSSupportingServiceConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(assc)
}
