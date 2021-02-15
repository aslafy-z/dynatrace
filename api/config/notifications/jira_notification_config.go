package notifications

import "github.com/dtcookie/gojson"

// JiraNotificationConfig Configuration of the Jira notification.
type JiraNotificationConfig struct {
	BaseNotificationConfig `json:",type=JIRA"`
	IssueType              string  `json:"issueType"`          // The type of the Jira issue to be created by this notification.
	Password               *string `json:"password,omitempty"` // The password for the Jira profile.
	ProjectKey             string  `json:"projectKey"`         // The project key of the Jira issue to be created by this notification.
	Summary                string  `json:"summary"`            // The summary of the Jira issue to be created by this notification.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	URL                    string  `json:"url"`                // The URL of the Jira API endpoint.
	Username               string  `json:"username"`           // The username of the Jira profile.
	Description            string  `json:"description"`        // The description of the Jira issue to be created by this notification.   You can use same placeholders as in issue summary.
}

// UnmarshalJSON provides custom JSON deserialization
func (jnc *JiraNotificationConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, jnc)
}

// MarshalJSON provides custom JSON serialization
func (jnc *JiraNotificationConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(jnc)
}
