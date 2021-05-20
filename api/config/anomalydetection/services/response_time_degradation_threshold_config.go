package services

import "github.com/dtcookie/gojson"

// ResponseTimeDegradationThresholdConfig Fixed thresholds for response time degradation detection.
//  Required if **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
type ResponseTimeDegradationThresholdConfig struct {
	LoadThreshold                            LoadThreshold `json:"loadThreshold"`                            // Minimal service load to detect response time degradation.   Response time degradation of services with smaller load won't trigger alerts.
	ResponseTimeThresholdMilliseconds        int32         `json:"responseTimeThresholdMilliseconds"`        // Response time during any 5-minute period to trigger an alert, in milliseconds.
	Sensitivity                              Sensitivity   `json:"sensitivity"`                              // Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers an alert.
	SlowestResponseTimeThresholdMilliseconds int32         `json:"slowestResponseTimeThresholdMilliseconds"` // Response time of the 10% slowest during any 5-minute period to trigger an alert, in milliseconds.
}

// UnmarshalJSON provides custom JSON deserialization
func (rtdtc *ResponseTimeDegradationThresholdConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, rtdtc)
}

// MarshalJSON provides custom JSON serialization
func (rtdtc *ResponseTimeDegradationThresholdConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(rtdtc)
}
