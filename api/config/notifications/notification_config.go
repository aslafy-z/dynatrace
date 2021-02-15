package notifications

import "github.com/dtcookie/gojson"

// NotificationConfig Configuration of a notification. The actual set of fields depends on the `type` of the notification.
// See the [Notifications API - JSON models](https://www.dynatrace.com/support/help/shortlink/api-config-notifications-models) help topic for example models of every notification type.
type NotificationConfig interface {
	Initialize(*BaseNotificationConfig) // Initialize duplicates the properties of the other AbstractConditionKey into this instance. It also serves as a unique method for structs implementing the interface this abstract class is derived from
	Implementors() []NotificationConfig // Implementors provides automatic JSON unmarshalling with a list of struct implementing (directly or indirectly) ConditionKey
	GetID() *string                     // The ID of the notification configuration.
}

// BaseNotificationConfig Configuration of a notification. The actual set of fields depends on the `type` of the notification.
// See the [Notifications API - JSON models](https://www.dynatrace.com/support/help/shortlink/api-config-notifications-models) help topic for example models of every notification type.
type BaseNotificationConfig struct {
	Name            string  `json:"name"`            // The name of the notification configuration.
	Type            Type    `json:"type"`            // Defines the actual set of fields depending on the value. See one of the following objects:  * `EMAIL` -> EmailNotificationConfig  * `PAGER_DUTY` -> PagerDutyNotificationConfig  * `WEBHOOK` -> WebHookNotificationConfig  * `SLACK` -> SlackNotificationConfig  * `HIPCHAT` -> HipChatNotificationConfig  * `VICTOROPS` -> VictorOpsNotificationConfig  * `SERVICE_NOW` -> ServiceNowNotificationConfig  * `XMATTERS` -> XMattersNotificationConfig  * `ANSIBLETOWER` -> AnsibleTowerNotificationConfig  * `OPS_GENIE` -> OpsGenieNotificationConfig  * `JIRA` -> JiraNotificationConfig  * `TRELLO` -> TrelloNotificationConfig
	Active          bool    `json:"active"`          // The configuration is enabled (`true`) or disabled (`false`).
	AlertingProfile string  `json:"alertingProfile"` // The ID of the associated alerting profile.
	ID              *string `json:"id,omitempty"`    // The ID of the notification configuration.
}

// GetID returns he ID of the notification configuration.
func (bnc *BaseNotificationConfig) GetID() *string {
	return bnc.ID
}

// Initialize duplicates the properties of the other BaseNotificationConfig into this instance
// It also serves as a unique method for structs implementing the interface this abstract class is derived from
func (bnc *BaseNotificationConfig) Initialize(other *BaseNotificationConfig) {
	bnc.Name = other.Name
	bnc.Type = other.Type
	bnc.Active = other.Active
	bnc.AlertingProfile = other.AlertingProfile
	bnc.ID = other.ID
}

// Implementors provides automatic JSON unmarshalling with a list of struct implementing (directly or indirectly) NotificationConfig
func (bnc *BaseNotificationConfig) Implementors() []NotificationConfig {
	return []NotificationConfig{
		new(WebHookNotificationConfig),
		new(EmailNotificationConfig),
		new(SlackNotificationConfig),
		new(PagerDutyNotificationConfig),
		new(TrelloNotificationConfig),
		new(VictorOpsNotificationConfig),
		new(XMattersNotificationConfig),
		new(OpsGenieNotificationConfig),
		new(JiraNotificationConfig),
		new(AnsibleTowerNotificationConfig),
		new(HipChatNotificationConfig),
		new(ServiceNowNotificationConfig),
	}
}

// UnmarshalJSON provides custom JSON deserialization
func (bnc *BaseNotificationConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, bnc)
}

// MarshalJSON provides custom JSON serialization
func (bnc *BaseNotificationConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(bnc)
}
