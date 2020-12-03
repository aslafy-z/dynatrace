package alertingprofiles

import "github.com/dtcookie/gojson"

// TagFilter A tag-based filter of monitored entities.
type TagFilter struct {
	Context Context `json:"context"`         // The origin of the tag, such as AWS or Cloud Foundry.  Custom tags use the `CONTEXTLESS` value.
	Key     string  `json:"key"`             // The key of the tag.  Custom tags have the tag value here.
	Value   *string `json:"value,omitempty"` // The value of the tag.   Not applicable to custom tags.
}

// UnmarshalJSON provides custom JSON deserialization
func (tf *TagFilter) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, tf)
}

// MarshalJSON provides custom JSON serialization
func (tf *TagFilter) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(tf)
}
