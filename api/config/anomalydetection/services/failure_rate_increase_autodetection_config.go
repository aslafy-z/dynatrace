package services

import "github.com/dtcookie/gojson"

// FailureRateIncreaseAutodetectionConfig Parameters of failure rate increase auto-detection. Required if **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.
// The absolute and relative thresholds **both** must exceed to trigger an alert.
// Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:
// Absolute: 1.5% + **1%** = 2.5%
// Relative: 1.5% + 1.5% * **50%** = 2.25%
type FailureRateIncreaseAutodetectionConfig struct {
	FailingServiceCallPercentageIncreaseAbsolute int32 `json:"failingServiceCallPercentageIncreaseAbsolute"` // Absolute increase of failing service calls to trigger an alert, %.
	FailingServiceCallPercentageIncreaseRelative int32 `json:"failingServiceCallPercentageIncreaseRelative"` // Relative increase of failing service calls to trigger an alert, %.
}

// UnmarshalJSON provides custom JSON deserialization
func (friac *FailureRateIncreaseAutodetectionConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, friac)
}

// MarshalJSON provides custom JSON serialization
func (friac *FailureRateIncreaseAutodetectionConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(friac)
}
