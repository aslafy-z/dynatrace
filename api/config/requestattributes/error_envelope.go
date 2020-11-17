package requestattributes

import "github.com/dtcookie/gojson"

// ErrorEnvelope has no documentation
type ErrorEnvelope struct {
	Error *Error `json:"error,omitempty"` // has no documentation
}

// UnmarshalJSON provides custom JSON deserialization
func (ee *ErrorEnvelope) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ee)
}

// MarshalJSON provides custom JSON serialization
func (ee *ErrorEnvelope) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ee)
}
