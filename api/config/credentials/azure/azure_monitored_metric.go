package azure

import "github.com/dtcookie/gojson"

// AzureMonitoredMetric A metric of supporting service to be monitored.
type AzureMonitoredMetric struct {
	Name       *string  `json:"name,omitempty"`       // The name of the metric of the supporting service.
	Dimensions []string `json:"dimensions,omitempty"` // A list of metric's dimensions names. It must include all the recommended dimensions.
}

// UnmarshalJSON provides custom JSON deserialization
func (amm *AzureMonitoredMetric) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, amm)
}

// MarshalJSON provides custom JSON serialization
func (amm *AzureMonitoredMetric) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(amm)
}
