package services

import "github.com/dtcookie/gojson"

// ResponseTimeDegradationDetectionConfig Configuration of response time degradation detection.
type ResponseTimeDegradationDetectionConfig struct {
	AutomaticDetection *ResponseTimeDegradationAutodetectionConfig `json:"automaticDetection,omitempty"` // Parameters of the response time degradation auto-detection. Required if the **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.  Violation of **any** criterion triggers an alert.
	DetectionMode      DetectionMode                               `json:"detectionMode"`                // How to detect response time degradation: automatically, or based on fixed thresholds, or do not detect.
	Thresholds         *ResponseTimeDegradationThresholdConfig     `json:"thresholds,omitempty"`         // Fixed thresholds for response time degradation detection.   Required if **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
}

// UnmarshalJSON provides custom JSON deserialization
func (rtddc *ResponseTimeDegradationDetectionConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, rtddc)
}

// MarshalJSON provides custom JSON serialization
func (rtddc *ResponseTimeDegradationDetectionConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(rtddc)
}
