package managementzones

import "github.com/dtcookie/gojson"

// DatabaseTopologyComparison Comparison for `DATABASE_TOPOLOGY` attributes.
type DatabaseTopologyComparison struct {
	BaseComparisonBasic `json:",type=DATABASE_TOPOLOGY"`
	Operator            DatabaseTopologyComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *DatabaseTopologyComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (dtc *DatabaseTopologyComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, dtc)
}

// MarshalJSON provides custom JSON serialization
func (dtc *DatabaseTopologyComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(dtc)
}
