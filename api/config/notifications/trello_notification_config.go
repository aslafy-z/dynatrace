package notifications

import "github.com/dtcookie/gojson"

// TrelloNotificationConfig Configuration of the Trello notification.
type TrelloNotificationConfig struct {
	BaseNotificationConfig `json:",type=TRELLO"`
	ResolvedListID         string  `json:"resolvedListId"`               // The Trello list to which the card of the resolved problem should be assigned.
	Text                   string  `json:"text"`                         // The text of the generated Trello card.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	ApplicationKey         string  `json:"applicationKey"`               // The application key for the Trello account.
	AuthorizationToken     *string `json:"authorizationToken,omitempty"` // The application token for the Trello account.
	BoardID                string  `json:"boardId"`                      // The Trello board to which the card should be assigned.
	Description            string  `json:"description"`                  // The description of the Trello card.   You can use same placeholders as in card text.
	ListID                 string  `json:"listId"`                       // The Trello list to which the card should be assigned.
}

// UnmarshalJSON provides custom JSON deserialization
func (tnc *TrelloNotificationConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, tnc)
}

// MarshalJSON provides custom JSON serialization
func (tnc *TrelloNotificationConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(tnc)
}
