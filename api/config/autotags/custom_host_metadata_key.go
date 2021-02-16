package autotags

import "github.com/dtcookie/gojson"

// CustomHostMetadataKey The key of the attribute, which need dynamic keys.
// Not applicable otherwise, as the attibute itself acts as a key.
type CustomHostMetadataKey struct {
	Key    string                      `json:"key"`    // The actual key of the custom metadata.
	Source CustomHostMetadataKeySource `json:"source"` // The source of the custom metadata.
}

// UnmarshalJSON provides custom JSON deserialization
func (chmk *CustomHostMetadataKey) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, chmk)
}

// MarshalJSON provides custom JSON serialization
func (chmk *CustomHostMetadataKey) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(chmk)
}
