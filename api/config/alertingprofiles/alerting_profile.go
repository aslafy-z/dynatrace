package alertingprofiles

import "github.com/dtcookie/gojson"

// AlertingProfile Configuration of an alerting profile.
type AlertingProfile struct {
	MzID             *string                       `json:"mzId,omitempty"`             // The ID of the management zone to which the alerting profile applies.
	Rules            []AlertingProfileSeverityRule `json:"rules,omitempty"`            // A list of severity rules.   The rules are evaluated from top to bottom. The first matching rule applies and further evaluation stops.  If you specify both severity rule and event filter, the AND logic applies.
	DisplayName      string                        `json:"displayName"`                // The name of the alerting profile, displayed in the UI.
	EventTypeFilters []AlertingEventTypeFilter     `json:"eventTypeFilters,omitempty"` // The list of event filters.  For all filters that are *negated* inside of these event filters, that is all "Predefined" as well as "Custom" (Title and/or Description) ones the AND logic applies. For all *non-negated* ones the OR logic applies. Between these two groups, negated and non-negated, the AND logic applies.  If you specify both severity rule and event filter, the AND logic applies.
	ID               *string                       `json:"id,omitempty"`               // The ID of the alerting profile.
	Metadata         *ConfigurationMetadata        `json:"metadata,omitempty"`         // Metadata useful for debugging
}

// UnmarshalJSON provides custom JSON deserialization
func (ap *AlertingProfile) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ap)
}

// MarshalJSON provides custom JSON serialization
func (ap *AlertingProfile) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ap)
}
