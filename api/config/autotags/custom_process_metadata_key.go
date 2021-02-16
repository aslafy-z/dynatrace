package autotags

import "github.com/dtcookie/gojson"

// CustomProcessMetadataKey The key of the attribute, which need dynamic keys.
// Not applicable otherwise, as the attibute itself acts as a key.
type CustomProcessMetadataKey struct {
	Key    string                         `json:"key"`    // The actual key of the custom metadata.
	Source CustomProcessMetadataKeySource `json:"source"` // The source of the custom metadata.
}

// UnmarshalJSON provides custom JSON deserialization
func (cpmk *CustomProcessMetadataKey) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, cpmk)
}

// MarshalJSON provides custom JSON serialization
func (cpmk *CustomProcessMetadataKey) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(cpmk)
}
