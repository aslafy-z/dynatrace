package requestattributes

import "github.com/dtcookie/gojson"

// ValueProcessing Process values as specified.
type ValueProcessing struct {
	ExtractSubstring    *ExtractSubstring `json:"extractSubstring,omitempty"`    // Preprocess by extracting a substring from the original value.
	SplitAt             *string           `json:"splitAt,omitempty"`             // Split (preprocessed) string values at this separator.
	Trim                bool              `json:"trim"`                          // Prune Whitespaces. Defaults to false.
	ValueCondition      *ValueCondition   `json:"valueCondition,omitempty"`      // IBM integration bus label node name condition for which the value is captured.
	ValueExtractorRegex *string           `json:"valueExtractorRegex,omitempty"` // Extract value from captured data per regex.
}

// UnmarshalJSON provides custom JSON deserialization
func (vp *ValueProcessing) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, vp)
}

// MarshalJSON provides custom JSON serialization
func (vp *ValueProcessing) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(vp)
}
