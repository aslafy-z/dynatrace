package managementzones

import "github.com/dtcookie/gojson"

// ServiceTypeComparison Comparison for `SERVICE_TYPE` attributes.
type ServiceTypeComparison struct {
	BaseComparisonBasic `json:",type=SERVICE_TYPE"`
	Operator            ServiceTypeComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *ServiceTypeComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (stc *ServiceTypeComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, stc)
}

// MarshalJSON provides custom JSON serialization
func (stc *ServiceTypeComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(stc)
}
