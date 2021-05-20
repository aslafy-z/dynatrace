package services

import "github.com/dtcookie/gojson"

// ServiceAnomalyDetectionConfig Dynatrace automatically detects service-related performance anomalies such as response time degradations and failure rate increases. Use these settings to configure detection sensitivity, set alert thresholds, or disable alerting for certain services.
type ServiceAnomalyDetectionConfig struct {
	LoadSpike               *LoadSpikeDetectionConfig              `json:"loadSpike,omitempty"`     // The configuration of load spikes detection.
	Metadata                *ConfigurationMetadata                 `json:"metadata,omitempty"`      // Metadata useful for debugging
	ResponseTimeDegradation ResponseTimeDegradationDetectionConfig `json:"responseTimeDegradation"` // Configuration of response time degradation detection.
	FailureRateIncrease     FailureRateIncreaseDetectionConfig     `json:"failureRateIncrease"`     // Configuration of failure rate increase detection.
	LoadDrop                *LoadDropDetectionConfig               `json:"loadDrop,omitempty"`      // The configuration of load drops detection.
}

// UnmarshalJSON provides custom JSON deserialization
func (sadc *ServiceAnomalyDetectionConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, sadc)
}

// MarshalJSON provides custom JSON serialization
func (sadc *ServiceAnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(sadc)
}
