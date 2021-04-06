package azure

import "github.com/dtcookie/gojson"

// AzureSupportingService A supporting service to be monitored.
type AzureSupportingService struct {
	MonitoredMetrics []AzureMonitoredMetric `json:"monitoredMetrics,omitempty"` // A list of metrics to be monitored for this service. It must include all the recommended metrics.
	Name             *string                `json:"name,omitempty"`             // The name of the supporting service.
}

// UnmarshalJSON provides custom JSON deserialization
func (ass *AzureSupportingService) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ass)
}

// MarshalJSON provides custom JSON serialization
func (ass *AzureSupportingService) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ass)
}
