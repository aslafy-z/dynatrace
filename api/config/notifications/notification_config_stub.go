package notifications

import "github.com/dtcookie/gojson"

// NotificationConfigStub The short representation of a notification.
type NotificationConfigStub struct {
	Description *string `json:"description,omitempty"` // A short description of the Dynatrace entity.
	ID          string  `json:"id"`                    // The ID of the Dynatrace entity.
	Name        *string `json:"name,omitempty"`        // The name of the Dynatrace entity.
	Type        *Type   `json:"type,omitempty"`        // The type of the notification.
}

// UnmarshalJSON provides custom JSON deserialization
func (ncs *NotificationConfigStub) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ncs)
}

// MarshalJSON provides custom JSON serialization
func (ncs *NotificationConfigStub) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ncs)
}
