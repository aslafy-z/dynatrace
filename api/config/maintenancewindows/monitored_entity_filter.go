package maintenancewindows

import "github.com/dtcookie/gojson"

// MonitoredEntityFilter A matching rule for Dynatrace entities.
type MonitoredEntityFilter struct {
	Type           *MonitoredEntityFilterType `json:"type,omitempty"`           // The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching.
	MzID           *string                    `json:"mzId,omitempty"`           // The ID of a management zone to which the matched entities must belong.
	TagCombination *TagCombination            `json:"tagCombination,omitempty"` // The logic that applies when several tags are specified: AND/OR.  If not set, the OR logic is used.
	Tags           []TagInfo                  `json:"tags"`                     // The tag you want to use for matching.  You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables.
}

// UnmarshalJSON provides custom JSON deserialization
func (mef *MonitoredEntityFilter) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, mef)
}

// MarshalJSON provides custom JSON serialization
func (mef *MonitoredEntityFilter) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(mef)
}
