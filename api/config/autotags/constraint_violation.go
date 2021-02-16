package autotags

import "github.com/dtcookie/gojson"

// ConstraintViolation has no documentation
type ConstraintViolation struct {
	Path              *string            `json:"path,omitempty"`              // has no documentation
	Location          *string            `json:"location,omitempty"`          // has no documentation
	Message           *string            `json:"message,omitempty"`           // has no documentation
	ParameterLocation *ParameterLocation `json:"parameterLocation,omitempty"` // has no documentation
}

// UnmarshalJSON provides custom JSON deserialization
func (cv *ConstraintViolation) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, cv)
}

// MarshalJSON provides custom JSON serialization
func (cv *ConstraintViolation) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(cv)
}
