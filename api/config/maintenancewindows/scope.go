package maintenancewindows

import "github.com/dtcookie/gojson"

// Scope The scope of the maintenance window.
//  The scope restricts the alert/problem detection suppression to certain Dynatrace entities. It can contain a list of entities and/or matching rules for dynamic formation of the scope.
//  If no scope is specified, the alert/problem detection suppression applies to the entire environment.
type Scope struct {
	Entities []string                `json:"entities"` // A list of Dynatrace entities (for example, hosts or services) to be included in the scope.  Allowed values are Dynatrace entity IDs.
	Matches  []MonitoredEntityFilter `json:"matches"`  // A list of matching rules for dynamic scope formation.  If several rules are set, the OR logic applies.
}

// UnmarshalJSON provides custom JSON deserialization
func (s *Scope) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, s)
}

// MarshalJSON provides custom JSON serialization
func (s *Scope) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(s)
}
