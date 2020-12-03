package alertingprofiles

import "github.com/dtcookie/gojson"

// AlertingEventTypeFilter Configuration of the event filter for the alerting profile.
// You have two mutually exclusive options:
// * Select an event type from the list of the predefined events. Specify it in the **predefinedEventFilter** field.
// * Set a rule for custom events. Specify it in the **customEventFilter** field.
type AlertingEventTypeFilter struct {
	CustomEventFilter     *AlertingCustomEventFilter     `json:"customEventFilter,omitempty"`     // Configuration of a custom event filter.  Filters custom events by title or description. If both specified, the AND logic applies.
	PredefinedEventFilter *AlertingPredefinedEventFilter `json:"predefinedEventFilter,omitempty"` // Configuration of a predefined event filter.
}

// UnmarshalJSON provides custom JSON deserialization
func (aetf *AlertingEventTypeFilter) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, aetf)
}

// MarshalJSON provides custom JSON serialization
func (aetf *AlertingEventTypeFilter) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(aetf)
}
