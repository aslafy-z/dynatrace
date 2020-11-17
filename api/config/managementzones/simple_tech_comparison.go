package managementzones

import "github.com/dtcookie/gojson"

// SimpleTechComparison Comparison for `SIMPLE_TECH` attributes.
type SimpleTechComparison struct {
	BaseComparisonBasic `json:",type=SIMPLE_TECH"`
	Operator            SimpleTechComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *SimpleTech                  `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (stc *SimpleTechComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, stc)
}

// MarshalJSON provides custom JSON serialization
func (stc *SimpleTechComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(stc)
}
