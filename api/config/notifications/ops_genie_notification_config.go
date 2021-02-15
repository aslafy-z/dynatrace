package notifications

import "github.com/dtcookie/gojson"

// OpsGenieNotificationConfig Configuration of the OpsGenie notification.
type OpsGenieNotificationConfig struct {
	BaseNotificationConfig `json:",type=OPS_GENIE"`
	APIKey                 *string `json:"apiKey,omitempty"` // The API key to access OpsGenie.
	Domain                 string  `json:"domain"`           // The region domain of the OpsGenie.
	Message                string  `json:"message"`          // The content of the message.  You can use the following placeholders:  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.
}

// UnmarshalJSON provides custom JSON deserialization
func (ognc *OpsGenieNotificationConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ognc)
}

// MarshalJSON provides custom JSON serialization
func (ognc *OpsGenieNotificationConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ognc)
}
