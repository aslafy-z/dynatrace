package autotags

import "github.com/dtcookie/gojson"

// TagComparison Comparison for `TAG` attributes.
type TagComparison struct {
	BaseComparisonBasic `json:",type=TAG"`
	Operator            TagComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *TagInfo              `json:"value,omitempty"` // Tag of a Dynatrace entity.
}

// UnmarshalJSON provides custom JSON deserialization
func (tc *TagComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, tc)
}

// MarshalJSON provides custom JSON serialization
func (tc *TagComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(tc)
}
