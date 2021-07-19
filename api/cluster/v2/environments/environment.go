package environments

import (
	"github.com/dtcookie/dynatrace/api/cluster/v2/environments/quota"
	"github.com/dtcookie/dynatrace/api/cluster/v2/environments/storage"
)

// Environment representas basic configuration for an environment
type Environment struct {
	Name    string            `json:"name"`              // The display name of the environment
	Trial   *bool             `json:"trial,omitempty"`   // Specifies whether the environment is a trial environment or a non-trial environment. Creating a trial environment is only possible if your license allows that. The default value is false (non-trial)
	State   *State            `json:"state,omitempty"`   // Indicates whether the environment is enabled or disabled. The default value is ENABLED
	Tags    []string          `json:"tags,omitempty"`    // A set of tags that are assigned to this environment. Every tag can have a maximum length of 100 characters
	Quotas  *quota.Settings   `json:"quotas,omitempty"`  // Environment level consumption and quotas information. Only returned if includeConsumptionInfo or includeUncachedConsumptionInfo param is true. If skipped when editing via PUT method then already set quotas will remain
	Storage *storage.Settings `json:"storage,omitempty"` // Environment level storage usage and limit information. Not returned if includeStorageInfo param is not true. If skipped when editing via PUT method then already set limits will remain
}

type State string

var States = struct {
	Enabled  State
	Disabled State
}{
	Enabled:  State("ENABLED"),
	Disabled: State("DISABLED"),
}
