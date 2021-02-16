package autotags

import "github.com/dtcookie/gojson"

// HypervisorTypeComparision Comparison for `HYPERVISOR_TYPE` attributes.
type HypervisorTypeComparision struct {
	BaseComparisonBasic `json:",type=HYPERVISOR_TYPE"`
	Operator            HypervisorTypeComparisionOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *HypervisorTypeComparisionValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (htc *HypervisorTypeComparision) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, htc)
}

// MarshalJSON provides custom JSON serialization
func (htc *HypervisorTypeComparision) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(htc)
}
