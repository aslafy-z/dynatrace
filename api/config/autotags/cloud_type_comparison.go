package autotags

import "github.com/dtcookie/gojson"

// CloudTypeComparison Comparison for `CLOUD_TYPE` attributes.
type CloudTypeComparison struct {
	BaseComparisonBasic `json:",type=CLOUD_TYPE"`
	Operator            CloudTypeComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *CloudTypeComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (ctc *CloudTypeComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ctc)
}

// MarshalJSON provides custom JSON serialization
func (ctc *CloudTypeComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ctc)
}
