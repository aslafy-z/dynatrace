package requestattributes

import "github.com/dtcookie/gojson"

// ExtractSubstring Preprocess by extracting a substring from the original value.
type ExtractSubstring struct {
	EndDelimiter *string  `json:"endDelimiter,omitempty"` // The end-delimiter string.   Required if the **position** value is `BETWEEN`. Otherwise not allowed.
	Position     Position `json:"position"`               // The position of the extracted string relative to delimiters.
	Delimiter    string   `json:"delimiter"`              // The delimiter string.
}

// UnmarshalJSON provides custom JSON deserialization
func (es *ExtractSubstring) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, es)
}

// MarshalJSON provides custom JSON serialization
func (es *ExtractSubstring) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(es)
}
