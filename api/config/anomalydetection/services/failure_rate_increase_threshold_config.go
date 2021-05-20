package services

import "github.com/dtcookie/gojson"

// FailureRateIncreaseThresholdConfig Fixed thresholds for failure rate increase detection.
//  Required if **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
type FailureRateIncreaseThresholdConfig struct {
	Sensitivity Sensitivity `json:"sensitivity"` // Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers alert.
	Threshold   int32       `json:"threshold"`   // Failure rate during any 5-minute period to trigger an alert, %.
}

// UnmarshalJSON provides custom JSON deserialization
func (fritc *FailureRateIncreaseThresholdConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, fritc)
}

// MarshalJSON provides custom JSON serialization
func (fritc *FailureRateIncreaseThresholdConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(fritc)
}
