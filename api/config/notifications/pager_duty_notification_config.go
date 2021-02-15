package notifications

import "github.com/dtcookie/gojson"

// PagerDutyNotificationConfig Configuration of the PagerDuty notification.
type PagerDutyNotificationConfig struct {
	BaseNotificationConfig `json:",type=PAGER_DUTY"`
	Account                string  `json:"account"`                 // The name of the PagerDuty account.
	ServiceAPIKey          *string `json:"serviceApiKey,omitempty"` // The API key to access PagerDuty.
	ServiceName            string  `json:"serviceName"`             // The name of the service.
}

// UnmarshalJSON provides custom JSON deserialization
func (pdnc *PagerDutyNotificationConfig) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, pdnc)
}

// MarshalJSON provides custom JSON serialization
func (pdnc *PagerDutyNotificationConfig) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(pdnc)
}
