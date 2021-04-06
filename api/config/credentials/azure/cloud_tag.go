package azure

import "github.com/dtcookie/gojson"

// CloudTag A cloud tag.
type CloudTag struct {
	Value *string `json:"value,omitempty"` // The value of the tag.   If set to `null`, then resources with any value of the tag are monitored.
	Name  *string `json:"name,omitempty"`  // The name of the tag.
}

// UnmarshalJSON provides custom JSON deserialization
func (ct *CloudTag) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ct)
}

// MarshalJSON provides custom JSON serialization
func (ct *CloudTag) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ct)
}
