package aws

import "github.com/dtcookie/gojson"

// AWSSupportingServiceMetric A metric of supporting service to be monitored.
type AWSSupportingServiceMetric struct {
	Name       string    `json:"name"`       // The name of the metric of the supporting service.
	Statistic  Statistic `json:"statistic"`  // The statistic (aggregation) to be used for the metric. AVG_MIN_MAX value is 3 statistics at once: AVERAGE, MINIMUM and MAXIMUM
	Dimensions []string  `json:"dimensions"` // A list of metric's dimensions names.
}

// UnmarshalJSON provides custom JSON deserialization
func (assm *AWSSupportingServiceMetric) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, assm)
}

// MarshalJSON provides custom JSON serialization
func (assm *AWSSupportingServiceMetric) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(assm)
}
