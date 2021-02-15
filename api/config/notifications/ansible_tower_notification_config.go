package notifications

import "github.com/dtcookie/gojson"

// AnsibleTowerNotificationConfig Configuration of the Ansible Tower notification.
type AnsibleTowerNotificationConfig struct {
	BaseNotificationConfig `json:",type=ANSIBLETOWER"`
	AcceptAnyCertificate   bool    `json:"acceptAnyCertificate"` // Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates.
	CustomMessage          string  `json:"customMessage"`        // The custom message of the notification.   This message will be displayed in the extra variables **Message** field of your job template.  You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	JobTemplateID          int32   `json:"jobTemplateID"`        // The ID of the target Ansible Tower job template.
	JobTemplateURL         string  `json:"jobTemplateURL"`       // The URL of the target Ansible Tower job template.
	Password               *string `json:"password,omitempty"`   // The password for the Ansible Tower account.
	Username               string  `json:"username"`             // The username of the Ansible Tower account.
}

// UnmarshalJSON provides custom JSON deserialization
func (atnc *AnsibleTowerNotificationConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, atnc)
}

// MarshalJSON provides custom JSON serialization
func (atnc *AnsibleTowerNotificationConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(atnc)
}
