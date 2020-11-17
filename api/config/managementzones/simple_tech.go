package managementzones

import "github.com/dtcookie/gojson"

// SimpleTech The value to compare to.
type SimpleTech struct {
	Type         *SimpleTechType `json:"type,omitempty"`         // Predefined technology, if technology is not predefined, then the verbatim type must be set
	VerbatimType *string         `json:"verbatimType,omitempty"` // Non-predefined technology, use for custom technologies.
}

// UnmarshalJSON provides custom JSON deserialization
func (st *SimpleTech) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, st)
}

// MarshalJSON provides custom JSON serialization
func (st *SimpleTech) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(st)
}
