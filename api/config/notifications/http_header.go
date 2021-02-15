package notifications

import "github.com/dtcookie/gojson"

// HTTPHeader The HTTP header.
type HTTPHeader struct {
	Name  string  `json:"name"`            // The name of the HTTP header.
	Value *string `json:"value,omitempty"` // The value of the HTTP header. May contain an empty value.   Required when creating a new notification.  For the **Authorization** header, GET requests return the `null` value.  If you want update a notification configuration with an **Authorization** header which you want to remain intact, set the **Authorization** header with the `null` value.
}

// UnmarshalJSON provides custom JSON deserialization
func (hh *HTTPHeader) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, hh)
}

// MarshalJSON provides custom JSON serialization
func (hh *HTTPHeader) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(hh)
}
