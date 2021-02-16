package autotags

import "github.com/dtcookie/gojson"

// ApplicationTypeComparison Comparison for `APPLICATION_TYPE` attributes.
type ApplicationTypeComparison struct {
	BaseComparisonBasic `json:",type=APPLICATION_TYPE"`
	Operator            ApplicationTypeComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *ApplicationTypeComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (atc *ApplicationTypeComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, atc)
}

// MarshalJSON provides custom JSON serialization
func (atc *ApplicationTypeComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(atc)
}
