package notifications

import "github.com/dtcookie/gojson"

// NotificationConfigStubListDto has no documentation
type NotificationConfigStubListDto struct {
	Values []NotificationConfigStub `json:"values,omitempty"` // has no documentation
}

// UnmarshalJSON provides custom JSON deserialization
func (ncsld *NotificationConfigStubListDto) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ncsld)
}

// MarshalJSON provides custom JSON serialization
func (ncsld *NotificationConfigStubListDto) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ncsld)
}
