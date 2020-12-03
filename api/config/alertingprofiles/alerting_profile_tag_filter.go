package alertingprofiles

import "github.com/dtcookie/gojson"

// AlertingProfileTagFilter Configuration of the tag filtering of the alerting profile.
type AlertingProfileTagFilter struct {
	IncludeMode IncludeMode `json:"includeMode"`          // The filtering mode:  * `INCLUDE_ANY`: The rule applies to monitored entities that have at least one of the specified tags. You can specify up to 100 tags.  * `INCLUDE_ALL`: The rule applies to monitored entities that have **all** of the specified tags. You can specify up to 10 tags.  * `NONE`: The rule applies to all monitored entities.
	TagFilters  []TagFilter `json:"tagFilters,omitempty"` // A list of required tags.
}

// UnmarshalJSON provides custom JSON deserialization
func (aptf *AlertingProfileTagFilter) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, aptf)
}

// MarshalJSON provides custom JSON serialization
func (aptf *AlertingProfileTagFilter) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(aptf)
}
