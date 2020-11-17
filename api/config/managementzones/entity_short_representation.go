package managementzones

import "github.com/dtcookie/gojson"

// EntityShortRepresentation The short representation of a Dynatrace entity.
type EntityShortRepresentation struct {
	Description *string `json:"description,omitempty"` // A short description of the Dynatrace entity.
	ID          string  `json:"id"`                    // The ID of the Dynatrace entity.
	Name        *string `json:"name,omitempty"`        // The name of the Dynatrace entity.
}

// UnmarshalJSON provides custom JSON deserialization
func (esr *EntityShortRepresentation) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, esr)
}

// MarshalJSON provides custom JSON serialization
func (esr *EntityShortRepresentation) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(esr)
}
