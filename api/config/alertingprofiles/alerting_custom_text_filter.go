package alertingprofiles

import "github.com/dtcookie/gojson"

// AlertingCustomTextFilter Configuration of a matching filter.
type AlertingCustomTextFilter struct {
	Enabled         bool     `json:"enabled"`         // The filter is enabled (`true`) or disabled (`false`).
	Negate          bool     `json:"negate"`          // Reverses the comparison **operator**. For example it turns the **begins with** into **does not begin with**.
	Operator        Operator `json:"operator"`        // Operator of the comparison.   You can reverse it by setting **negate** to `true`.
	Value           string   `json:"value"`           // The value to compare to.
	CaseInsensitive bool     `json:"caseInsensitive"` // The condition is case sensitive (`false`) or case insensitive (`true`).   If not set, then `false` is used, making the condition case sensitive.
}

// UnmarshalJSON provides custom JSON deserialization
func (actf *AlertingCustomTextFilter) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, actf)
}

// MarshalJSON provides custom JSON serialization
func (actf *AlertingCustomTextFilter) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(actf)
}
