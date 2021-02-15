package notifications

// Type Defines the actual set of fields depending on the value. See one of the following objects:
// * `EMAIL` -> EmailNotificationConfig
// * `PAGER_DUTY` -> PagerDutyNotificationConfig
// * `WEBHOOK` -> WebHookNotificationConfig
// * `SLACK` -> SlackNotificationConfig
// * `HIPCHAT` -> HipChatNotificationConfig
// * `VICTOROPS` -> VictorOpsNotificationConfig
// * `SERVICE_NOW` -> ServiceNowNotificationConfig
// * `XMATTERS` -> XMattersNotificationConfig
// * `ANSIBLETOWER` -> AnsibleTowerNotificationConfig
// * `OPS_GENIE` -> OpsGenieNotificationConfig
// * `JIRA` -> JiraNotificationConfig
// * `TRELLO` -> TrelloNotificationConfig
type Type string

// Types offers the known enum values
var Types = struct {
	Ansibletower Type
	Email        Type
	Hipchat      Type
	Jira         Type
	OpsGenie     Type
	PagerDuty    Type
	ServiceNow   Type
	Slack        Type
	Trello       Type
	Victorops    Type
	Webhook      Type
	Xmatters     Type
}{
	"ANSIBLETOWER",
	"EMAIL",
	"HIPCHAT",
	"JIRA",
	"OPS_GENIE",
	"PAGER_DUTY",
	"SERVICE_NOW",
	"SLACK",
	"TRELLO",
	"VICTOROPS",
	"WEBHOOK",
	"XMATTERS",
}
