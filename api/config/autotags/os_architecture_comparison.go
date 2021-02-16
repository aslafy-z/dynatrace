package autotags

import "github.com/dtcookie/gojson"

// OSArchitectureComparison Comparison for `OS_ARCHITECTURE` attributes.
type OSArchitectureComparison struct {
	BaseComparisonBasic `json:",type=OS_ARCHITECTURE"`
	Operator            OSArchitectureComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *OSArchitectureComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (oac *OSArchitectureComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, oac)
}

// MarshalJSON provides custom JSON serialization
func (oac *OSArchitectureComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(oac)
}
