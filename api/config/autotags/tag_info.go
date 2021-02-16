package autotags

import "github.com/dtcookie/gojson"

// TagInfo Tag of a Dynatrace entity.
type TagInfo struct {
	Context Context `json:"context"`         // The origin of the tag, such as AWS or Cloud Foundry.   Custom tags use the `CONTEXTLESS` value.
	Key     string  `json:"key"`             // The key of the tag.   Custom tags have the tag value here.
	Value   *string `json:"value,omitempty"` // The value of the tag.   Not applicable to custom tags.
}

// UnmarshalJSON provides custom JSON deserialization
func (ti *TagInfo) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ti)
}

// MarshalJSON provides custom JSON serialization
func (ti *TagInfo) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ti)
}
