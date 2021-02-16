package autotags

import "github.com/dtcookie/gojson"

// SimpleHostTechComparison Comparison for `SIMPLE_HOST_TECH` attributes.
type SimpleHostTechComparison struct {
	BaseComparisonBasic `json:",type=SIMPLE_HOST_TECH"`
	Operator            SimpleHostTechComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *SimpleHostTech                  `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (shtc *SimpleHostTechComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, shtc)
}

// MarshalJSON provides custom JSON serialization
func (shtc *SimpleHostTechComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(shtc)
}
