package managementzones

import "github.com/dtcookie/gojson"

// OSTypeComparison Comparison for `OS_TYPE` attributes.
type OSTypeComparison struct {
	BaseComparisonBasic `json:",type=OS_TYPE"`
	Operator            OSTypeComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *OSTypeComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (otc *OSTypeComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, otc)
}

// MarshalJSON provides custom JSON serialization
func (otc *OSTypeComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(otc)
}
