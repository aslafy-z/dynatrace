package maintenancewindows

import "github.com/dtcookie/gojson"

// MaintenanceWindow Configuration of a maintenance window.
type MaintenanceWindow struct {
	Schedule    Schedule               `json:"schedule"`           // The schedule of the maintenance window.
	Scope       *Scope                 `json:"scope,omitempty"`    // The scope of the maintenance window.   The scope restricts the alert/problem detection suppression to certain Dynatrace entities. It can contain a list of entities and/or matching rules for dynamic formation of the scope.   If no scope is specified, the alert/problem detection suppression applies to the entire environment.
	Suppression Suppression            `json:"suppression"`        // The type of suppression of alerting and problem detection during the maintenance.
	Type        MaintenanceWindowType  `json:"type"`               // The type of the maintenance: planned or unplanned.
	Description string                 `json:"description"`        // A short description of the maintenance purpose.
	ID          *string                `json:"id,omitempty"`       // The ID of the maintenance window.
	Metadata    *ConfigurationMetadata `json:"metadata,omitempty"` // Metadata useful for debugging
	Name        string                 `json:"name"`               // The name of the maintenance window, displayed in the UI.
}

// UnmarshalJSON provides custom JSON deserialization
func (mw *MaintenanceWindow) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, mw)
}

// MarshalJSON provides custom JSON serialization
func (mw *MaintenanceWindow) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(mw)
}
