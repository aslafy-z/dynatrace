package alertingprofiles

import "github.com/dtcookie/gojson"

// AlertingProfileSeverityRule A severity rule of the alerting profile.
//  A severity rule defines the level of severity that must be met before an alert is sent our for a detected problem. Additionally it restricts the alerting to certain monitored entities.
type AlertingProfileSeverityRule struct {
	TagFilter      AlertingProfileTagFilter `json:"tagFilter"`      // Configuration of the tag filtering of the alerting profile.
	DelayInMinutes int32                    `json:"delayInMinutes"` // Send a notification if a problem remains open longer than *X* minutes.
	SeverityLevel  SeverityLevel            `json:"severityLevel"`  // The severity level to trigger the alert.
}

// UnmarshalJSON provides custom JSON deserialization
func (apsr *AlertingProfileSeverityRule) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, apsr)
}

// MarshalJSON provides custom JSON serialization
func (apsr *AlertingProfileSeverityRule) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(apsr)
}
