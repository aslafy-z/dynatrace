package autotags

import "github.com/dtcookie/gojson"

// EntityIDComparison Comparison for `ENTITY_ID` attributes.
type EntityIDComparison struct {
	BaseComparisonBasic `json:",type=ENTITY_ID"`
	Value               *string                    `json:"value,omitempty"` // The value to compare to.
	Operator            EntityIDComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
}

// UnmarshalJSON provides custom JSON deserialization
func (eic *EntityIDComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, eic)
}

// MarshalJSON provides custom JSON serialization
func (eic *EntityIDComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(eic)
}
