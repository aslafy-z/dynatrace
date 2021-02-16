package autotags

import "github.com/dtcookie/gojson"

// StringConditionKey The key for dynamic attributes of the `STRING` type.
type StringConditionKey struct {
	BaseConditionKey `json:",type=STRING"`
	DynamicKey       string `json:"dynamicKey"` // The key of the attribute, which need dynamic keys.  Not applicable otherwise, as the attibute itself acts as a key.
}

// UnmarshalJSON provides custom JSON deserialization
func (sck *StringConditionKey) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, sck)
}

// MarshalJSON provides custom JSON serialization
func (sck *StringConditionKey) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(sck)
}
