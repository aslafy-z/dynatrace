package managementzones

import "github.com/dtcookie/gojson"

// PaasTypeComparison Comparison for `PAAS_TYPE` attributes.
type PaasTypeComparison struct {
	BaseComparisonBasic `json:",type=PAAS_TYPE"`
	Operator            PaasTypeComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *PaasTypeComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (ptc *PaasTypeComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ptc)
}

// MarshalJSON provides custom JSON serialization
func (ptc *PaasTypeComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ptc)
}
