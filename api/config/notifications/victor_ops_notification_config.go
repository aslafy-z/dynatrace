package notifications

import "github.com/dtcookie/gojson"

// VictorOpsNotificationConfig Configuration of the VictorOps notification.
type VictorOpsNotificationConfig struct {
	BaseNotificationConfig `json:",type=VICTOROPS"`
	APIKey                 *string `json:"apiKey,omitempty"` // The API key for the target VictorOps account.
	Message                string  `json:"message"`          // The content of the message.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.
	RoutingKey             string  `json:"routingKey"`       // The routing key, defining the group to be notified.
}

// UnmarshalJSON provides custom JSON deserialization
func (vonc *VictorOpsNotificationConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, vonc)
}

// MarshalJSON provides custom JSON serialization
func (vonc *VictorOpsNotificationConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(vonc)
}
