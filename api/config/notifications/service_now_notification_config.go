package notifications

import "github.com/dtcookie/gojson"

// ServiceNowNotificationConfig Configuration of the ServiceNow notification.
type ServiceNowNotificationConfig struct {
	BaseNotificationConfig `json:",type=SERVICE_NOW"`
	SendEvents             bool    `json:"sendEvents"`             // Send events into ServiceNow ITOM (`true`).
	SendIncidents          bool    `json:"sendIncidents"`          // Send incidents into ServiceNow ITSM (`true`).
	URL                    *string `json:"url,omitempty"`          // The URL of the on-premise ServiceNow installation.   This field is mutually exclusive with the **instanceName** field. You can only use one of them.
	Username               string  `json:"username"`               // The username of the ServiceNow account.   Make sure that your user account has the `rest_service`, `web_request_admin`, and `x_dynat_ruxit.Integration` roles.
	InstanceName           *string `json:"instanceName,omitempty"` // The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL.   This field is mutually exclusive with the **url** field. You can only use one of them.
	Message                string  `json:"message"`                // The content of the ServiceNow description.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	Password               *string `json:"password,omitempty"`     // The username to the ServiceNow account
}

// UnmarshalJSON provides custom JSON deserialization
func (snnc *ServiceNowNotificationConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, snnc)
}

// MarshalJSON provides custom JSON serialization
func (snnc *ServiceNowNotificationConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(snnc)
}
