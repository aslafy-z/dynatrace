package managementzones

import "github.com/dtcookie/gojson"

// ManagementZoneRule The rule of the management zone usage. It defines how the management zone applies.
// Each rule is evaluated independently of all other rules.
type ManagementZoneRule struct {
	Conditions       []EntityRuleEngineCondition `json:"conditions"`                 // A list of matching rules for the management zone.  The management zone applies only if **all** conditions are fulfilled.
	Enabled          bool                        `json:"enabled"`                    // The rule is enabled (`true`) or disabled (`false`).
	PropagationTypes []PropagationType           `json:"propagationTypes,omitempty"` // How to apply the management zone to underlying entities:  * `SERVICE_TO_HOST_LIKE`: Apply to underlying hosts of matching services.  * `SERVICE_TO_PROCESS_GROUP_LIKE`: Apply to underlying process groups of matching services.  * `PROCESS_GROUP_TO_HOST`: Apply to underlying hosts of matching process groups.  * `PROCESS_GROUP_TO_SERVICE`: Apply to all services provided by matching process groups.  * `HOST_TO_PROCESS_GROUP_INSTANCE`: Apply to processes running on matching hosts.  * `CUSTOM_DEVICE_GROUP_TO_CUSTOM_DEVICE`: Apply to custom devices in matching custom device groups.  * `AZURE_TO_PG`: Apply to process groups connected to matching Azure entities.  * `AZURE_TO_SERVICE`: Apply to services provided by matching Azure entities.
	Type             ManagementZoneRuleType      `json:"type"`                       // The type of Dynatrace entities the management zone can be applied to.
}

// UnmarshalJSON provides custom JSON deserialization
func (mzr *ManagementZoneRule) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, mzr)
}

// MarshalJSON provides custom JSON serialization
func (mzr *ManagementZoneRule) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(mzr)
}
