package alertingprofiles

import "github.com/dtcookie/gojson"

// AlertingPredefinedEventFilter Configuration of a predefined event filter.
type AlertingPredefinedEventFilter struct {
	EventType EventType `json:"eventType"` // The type of the predefined event.
	Negate    bool      `json:"negate"`    // The alert triggers when the problem of specified severity arises while the specified event **is** happening (`false`) or while the specified event is **not** happening (`true`).   For example, if you chose the Slowdown (`PERFORMANCE`) severity and Unexpected high traffic (`APPLICATION_UNEXPECTED_HIGH_LOAD`) event with **negate** set to `true`, the alerting profile will trigger only when the slowdown problem is raised while there is no unexpected high traffic event.  Consider the following use case as an example. The Slowdown (`PERFORMANCE`) severity rule is set. Depending on the configuration of the event filter (Unexpected high traffic (`APPLICATION_UNEXPECTED_HIGH_LOAD`) event is used as an example), the behavior of the alerting profile is one of the following:* **negate** is set to `false`: The alert triggers when the slowdown problem is raised while unexpected high traffic event is happening.  * **negate** is set to `true`: The alert triggers when the slowdown problem is raised while there is no unexpected high traffic event.  * no event rule is set: The alert triggers when the slowdown problem is raised, regardless of any events.
}

// UnmarshalJSON provides custom JSON deserialization
func (apef *AlertingPredefinedEventFilter) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, apef)
}

// MarshalJSON provides custom JSON serialization
func (apef *AlertingPredefinedEventFilter) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(apef)
}
