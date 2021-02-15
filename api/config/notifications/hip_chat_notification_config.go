package notifications

import "github.com/dtcookie/gojson"

// HipChatNotificationConfig Configuration of the HipChat notification.
type HipChatNotificationConfig struct {
	BaseNotificationConfig `json:",type=HIPCHAT"`
	Message                string  `json:"message"`       // The content of the notification message.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	URL                    *string `json:"url,omitempty"` // The URL of the HipChat WebHook.  This is confidential information, therefore GET requests return this field with the `null` value, and it is optional for PUT requests.
}

// UnmarshalJSON provides custom JSON deserialization
func (hcnc *HipChatNotificationConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, hcnc)
}

// MarshalJSON provides custom JSON serialization
func (hcnc *HipChatNotificationConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(hcnc)
}
