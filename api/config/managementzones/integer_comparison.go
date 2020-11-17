package managementzones

import "github.com/dtcookie/gojson"

// IntegerComparison Comparison for `INTEGER` attributes.
type IntegerComparison struct {
	BaseComparisonBasic `json:",type=INTEGER"`
	Operator            IntegerComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *int32                    `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (ic *IntegerComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ic)
}

// MarshalJSON provides custom JSON serialization
func (ic *IntegerComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ic)
}
