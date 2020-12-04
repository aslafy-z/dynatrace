package services

import "github.com/dtcookie/gojson"

// LoadSpikeDetectionConfig The configuration of load spikes detection.
type LoadSpikeDetectionConfig struct {
	Enabled                           bool   `json:"enabled"`                                     // The detection is enabled (`true`) or disabled (`false`).
	LoadSpikePercent                  *int32 `json:"loadSpikePercent,omitempty"`                  // Alert if the observed load is more than *X* % of the expected value.
	MinAbnormalStateDurationInMinutes *int32 `json:"minAbnormalStateDurationInMinutes,omitempty"` // Alert if the service stays in abnormal state for at least *X* minutes.
}

// UnmarshalJSON provides custom JSON deserialization
func (lsdc *LoadSpikeDetectionConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, lsdc)
}

// MarshalJSON provides custom JSON serialization
func (lsdc *LoadSpikeDetectionConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(lsdc)
}
