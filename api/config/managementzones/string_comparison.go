package managementzones

import "github.com/dtcookie/gojson"

// StringComparison Comparison for `STRING` attributes.
type StringComparison struct {
	BaseComparisonBasic `json:",type=STRING"`
	CaseSensitive       *bool                    `json:"caseSensitive,omitempty"` // The comparison is case-sensitive (`true`) or insensitive (`false`).
	Operator            StringComparisonOperator `json:"operator"`                // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *string                  `json:"value,omitempty"`         // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (sc *StringComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, sc)
}

// MarshalJSON provides custom JSON serialization
func (sc *StringComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(sc)
}
