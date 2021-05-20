package services

import "github.com/dtcookie/gojson"

// ResponseTimeDegradationAutodetectionConfig Parameters of the response time degradation auto-detection. Required if the **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.
// Violation of **any** criterion triggers an alert.
type ResponseTimeDegradationAutodetectionConfig struct {
	LoadThreshold                              LoadThreshold `json:"loadThreshold"`                              // Minimal service load to detect response time degradation.   Response time degradation of services with smaller load won't trigger alerts.
	ResponseTimeDegradationMilliseconds        int32         `json:"responseTimeDegradationMilliseconds"`        // Alert if the response time degrades by more than *X* milliseconds.
	ResponseTimeDegradationPercent             int32         `json:"responseTimeDegradationPercent"`             // Alert if the response time degrades by more than *X* %.
	SlowestResponseTimeDegradationMilliseconds int32         `json:"slowestResponseTimeDegradationMilliseconds"` // Alert if the response time of the slowest 10% degrades by more than *X* milliseconds.
	SlowestResponseTimeDegradationPercent      int32         `json:"slowestResponseTimeDegradationPercent"`      // Alert if the response time of the slowest 10% degrades by more than *X* %.
}

// UnmarshalJSON provides custom JSON deserialization
func (rtdac *ResponseTimeDegradationAutodetectionConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, rtdac)
}

// MarshalJSON provides custom JSON serialization
func (rtdac *ResponseTimeDegradationAutodetectionConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(rtdac)
}
