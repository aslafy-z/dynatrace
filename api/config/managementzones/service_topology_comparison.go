package managementzones

import "github.com/dtcookie/gojson"

// ServiceTopologyComparison Comparison for `SERVICE_TOPOLOGY` attributes.
type ServiceTopologyComparison struct {
	BaseComparisonBasic `json:",type=SERVICE_TOPOLOGY"`
	Value               *ServiceTopologyComparisonValue   `json:"value,omitempty"` // The value to compare to.
	Operator            ServiceTopologyComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
}

// UnmarshalJSON provides custom JSON deserialization
func (stc *ServiceTopologyComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, stc)
}

// MarshalJSON provides custom JSON serialization
func (stc *ServiceTopologyComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(stc)
}
