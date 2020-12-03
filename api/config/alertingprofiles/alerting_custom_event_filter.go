package alertingprofiles

import "github.com/dtcookie/gojson"

// AlertingCustomEventFilter Configuration of a custom event filter.
// Filters custom events by title or description. If both specified, the AND logic applies.
type AlertingCustomEventFilter struct {
	CustomDescriptionFilter *AlertingCustomTextFilter `json:"customDescriptionFilter,omitempty"` // Configuration of a matching filter.
	CustomTitleFilter       *AlertingCustomTextFilter `json:"customTitleFilter,omitempty"`       // Configuration of a matching filter.
}

// UnmarshalJSON provides custom JSON deserialization
func (acef *AlertingCustomEventFilter) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, acef)
}

// MarshalJSON provides custom JSON serialization
func (acef *AlertingCustomEventFilter) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(acef)
}


