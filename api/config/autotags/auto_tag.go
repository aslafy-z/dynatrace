package autotags

import "github.com/dtcookie/gojson"

// AutoTag Configuration of an auto-tag. It defines the conditions of tag usage and the tag value.
type AutoTag struct {
	Rules    []AutoTagRule          `json:"rules,omitempty"`    // The list of rules for tag usage.  When there are multiple rules, the OR logic applies.
	ID       *string                `json:"id,omitempty"`       // The ID of the auto-tag.
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"` // Metadata useful for debugging
	Name     string                 `json:"name"`               // The name of the auto-tag, which is applied to entities.  Additionally you can specify a **valueFormat** in the tag rule. In that case the tag is used in the `name:valueFormat` format.  For example you can extend the `Infrastructure` tag to `Infrastructure:Windows` and `Infrastructure:Linux`.
}

// UnmarshalJSON provides custom JSON deserialization
func (at *AutoTag) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, at)
}

// MarshalJSON provides custom JSON serialization
func (at *AutoTag) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(at)
}
