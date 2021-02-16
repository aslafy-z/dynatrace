package autotags

import "github.com/dtcookie/gojson"

// CustomApplicationTypeComparison Comparison for `CUSTOM_APPLICATION_TYPE` attributes.
type CustomApplicationTypeComparison struct {
	BaseComparisonBasic `json:",type=CUSTOM_APPLICATION_TYPE"`
	Operator            CustomApplicationTypeComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *CustomApplicationTypeComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (catc *CustomApplicationTypeComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, catc)
}

// MarshalJSON provides custom JSON serialization
func (catc *CustomApplicationTypeComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(catc)
}
