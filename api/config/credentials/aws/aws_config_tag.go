package aws

import "github.com/dtcookie/gojson"

// AWSConfigTag An AWS tag of the resource to be monitored.
type AWSConfigTag struct {
	Name  string `json:"name"`  // The key of the AWS tag.
	Value string `json:"value"` // The value of the AWS tag.
}

// UnmarshalJSON provides custom JSON deserialization
func (act *AWSConfigTag) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, act)
}

// MarshalJSON provides custom JSON serialization
func (act *AWSConfigTag) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(act)
}
