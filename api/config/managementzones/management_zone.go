package managementzones

import "github.com/dtcookie/gojson"

// ManagementZone The configuration of the management zone. It defines how the management zone applies.
type ManagementZone struct {
	ID       *string                `json:"id,omitempty"`       // The ID of the management zone.
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"` // Metadata useful for debugging
	Name     string                 `json:"name"`               // The name of the management zone.
	Rules    []ManagementZoneRule   `json:"rules,omitempty"`    // A list of rules for management zone usage.  Each rule is evaluated independently of all other rules.
}

// UnmarshalJSON provides custom JSON deserialization
func (mz *ManagementZone) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, mz)
}

// MarshalJSON provides custom JSON serialization
func (mz *ManagementZone) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(mz)
}
