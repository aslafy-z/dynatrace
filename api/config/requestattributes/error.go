package requestattributes

import "github.com/dtcookie/gojson"

// Error has no documentation
type Error struct {
	Code                 *int32                `json:"code,omitempty"`                 // has no documentation
	ConstraintViolations []ConstraintViolation `json:"constraintViolations,omitempty"` // has no documentation
	Message              *string               `json:"message,omitempty"`              // has no documentation
}

// UnmarshalJSON provides custom JSON deserialization
func (e *Error) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, e)
}

// MarshalJSON provides custom JSON serialization
func (e *Error) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(e)
}
