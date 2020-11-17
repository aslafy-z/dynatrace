package managementzones

import "github.com/dtcookie/gojson"

// SyntheticEngineTypeComparison Comparison for `SYNTHETIC_ENGINE_TYPE` attributes.
type SyntheticEngineTypeComparison struct {
	BaseComparisonBasic `json:",type=SYNTHETIC_ENGINE_TYPE"`
	Operator            SyntheticEngineTypeComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *SyntheticEngineTypeComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (setc *SyntheticEngineTypeComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, setc)
}

// MarshalJSON provides custom JSON serialization
func (setc *SyntheticEngineTypeComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(setc)
}
