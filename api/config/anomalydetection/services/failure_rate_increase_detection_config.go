package services

import "github.com/dtcookie/gojson"

// FailureRateIncreaseDetectionConfig Configuration of failure rate increase detection.
type FailureRateIncreaseDetectionConfig struct {
	Thresholds         *FailureRateIncreaseThresholdConfig     `json:"thresholds,omitempty"`         // Fixed thresholds for failure rate increase detection.   Required if **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
	AutomaticDetection *FailureRateIncreaseAutodetectionConfig `json:"automaticDetection,omitempty"` // Parameters of failure rate increase auto-detection. Required if **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.  The absolute and relative thresholds **both** must exceed to trigger an alert.  Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
	DetectionMode      DetectionMode                           `json:"detectionMode"`                // How to detect failure rate increase: automatically, or based on fixed thresholds, or do not detect.
}

// UnmarshalJSON provides custom JSON deserialization
func (fridc *FailureRateIncreaseDetectionConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, fridc)
}

// MarshalJSON provides custom JSON serialization
func (fridc *FailureRateIncreaseDetectionConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(fridc)
}
