package customservices

import (
	"encoding/json"

	api "github.com/dtcookie/trace/api/config"
)

// CustomService has no documentation
type CustomService struct {
	Metadata *api.ConfigurationMetadata `json:"metadata,omitempty"`

	ID                  string              `json:"id,omitempty"`                  // The ID of the custom service (UUID)
	Name                string              `json:"name"`                          // The name of the custom service, displayed in the UI
	Order               string              `json:"order,omitempty"`               // The order string. Sorting custom services alphabetically by their order string determines their relative ordering. Typically this is managed by Dynatrace internally and will not be present in GET responses
	Enabled             bool                `json:"enabled"`                       // Custom service enabled/disabled
	Rules               []DetectionRule     `json:"rules,omitempty"`               // The list of rules defining the custom service
	QueueEntryPoint     bool                `json:"queueEntryPoint"`               // The queue entry point flag. Set to `true` for custom messaging services
	QueueEntryPointType QueueEntryPointType `json:"queueEntryPointType,omitempty"` // The queue entry point type
	ProcessGroups       []string            `json:"processGroups,omitempty"`       // The list of process groups the custom service should belong to
}

// String provides a string representation in JSON format for debugging purposes
func (cs CustomService) String() string {
	var result string
	if bytes, err := json.MarshalIndent(cs, "", "  "); err != nil {
		result = err.Error()
	} else {
		result = string(bytes)
	}
	return result
}
