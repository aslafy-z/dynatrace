package services

import "github.com/dtcookie/gojson"

// LoadDropDetectionConfig The configuration of load drops detection.
type LoadDropDetectionConfig struct {
	Enabled                           bool   `json:"enabled"`                                     // The detection is enabled (`true`) or disabled (`false`).
	LoadDropPercent                   *int32 `json:"loadDropPercent,omitempty"`                   // Alert if the observed load is less than *X* % of the expected value.
	MinAbnormalStateDurationInMinutes *int32 `json:"minAbnormalStateDurationInMinutes,omitempty"` // Alert if the service stays in abnormal state for at least *X* minutes.
}

// UnmarshalJSON provides custom JSON deserialization
func (lddc *LoadDropDetectionConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, lddc)
}

// MarshalJSON provides custom JSON serialization
func (lddc *LoadDropDetectionConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(lddc)
}
