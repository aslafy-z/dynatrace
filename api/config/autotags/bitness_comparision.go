package autotags

import "github.com/dtcookie/gojson"

// BitnessComparision Comparison for `BITNESS` attributes.
type BitnessComparision struct {
	BaseComparisonBasic `json:",type=BITNESS"`
	Operator            BitnessComparisionOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *BitnessComparisionValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (bc *BitnessComparision) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, bc)
}

// MarshalJSON provides custom JSON serialization
func (bc *BitnessComparision) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(bc)
}
