package managementzones

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// ConfigurationMetadata Metadata useful for debugging
type ConfigurationMetadata struct {
	ClusterVersion               *string  `json:"clusterVersion,omitempty"`               // Dynatrace server version.
	ConfigurationVersions        []int64  `json:"configurationVersions,omitempty"`        // A Sorted list of the version numbers of the configuration.
	CurrentConfigurationVersions []string `json:"currentConfigurationVersions,omitempty"` // A Sorted list of string version numbers of the configuration.
}

func (cm *ConfigurationMetadata) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cluster_version": {
			Type:        hcl.TypeString,
			Description: "Dynatrace server version",
			Optional:    true,
		},
		"configuration_versions": {
			Type:        hcl.TypeList,
			Description: "A Sorted list of the version numbers of the configuration",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeInt},
		},
		"current_configuration_versions": {
			Type:        hcl.TypeList,
			Description: "A Sorted list of the version numbers of the configuration",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (cm *ConfigurationMetadata) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if cm.ClusterVersion != nil && len(*cm.ClusterVersion) > 0 {
		result["cluster_version"] = *cm.ClusterVersion
	}
	if len(cm.ConfigurationVersions) > 0 {
		result["configuration_versions"] = cm.ConfigurationVersions
	}
	if len(cm.CurrentConfigurationVersions) > 0 {
		result["current_configuration_versions"] = cm.CurrentConfigurationVersions
	}
	return result, nil
}

func (cm *ConfigurationMetadata) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("cluster_version"); ok {
		cm.ClusterVersion = opt.NewString(value.(string))
	}
	if _, ok := decoder.GetOk("configuration_versions.#"); ok {
		cm.ConfigurationVersions = []int64{}
		if entries, ok := decoder.GetOk("configuration_versions"); ok {
			for _, entry := range entries.([]interface{}) {
				cm.ConfigurationVersions = append(cm.ConfigurationVersions, int64(entry.(int)))
			}
		}
	}
	if _, ok := decoder.GetOk("current_configuration_versions.#"); ok {
		cm.CurrentConfigurationVersions = []string{}
		if entries, ok := decoder.GetOk("current_configuration_versions"); ok {
			for _, entry := range entries.([]interface{}) {
				cm.CurrentConfigurationVersions = append(cm.CurrentConfigurationVersions, entry.(string))
			}
		}
	}
	return nil
}
