package requestattributes

import "github.com/dtcookie/gojson"

// ValueCondition IBM integration bus label node name condition for which the value is captured.
type ValueCondition struct {
	Negate   bool     `json:"negate"`   // Negate the comparison.
	Operator Operator `json:"operator"` // Operator comparing the extracted value to the comparison value.
	Value    string   `json:"value"`    // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (vc *ValueCondition) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, vc)
}

// MarshalJSON provides custom JSON serialization
func (vc *ValueCondition) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(vc)
}
