package requestattributes

import "github.com/dtcookie/gojson"

// ScopeConditions Conditions for data capturing.
type ScopeConditions struct {
	HostGroup         *string            `json:"hostGroup,omitempty"`         // Only applies to this host group.
	ProcessGroup      *string            `json:"processGroup,omitempty"`      // Only applies to this process group. Note that this can't be transferred between different clusters or environments.
	ServiceTechnology *ServiceTechnology `json:"serviceTechnology,omitempty"` // Only applies to this service technology.
	TagOfProcessGroup *string            `json:"tagOfProcessGroup,omitempty"` // Only apply to process groups matching this tag.
}

// UnmarshalJSON provides custom JSON deserialization
func (sc *ScopeConditions) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, sc)
}

// MarshalJSON provides custom JSON serialization
func (sc *ScopeConditions) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(sc)
}
