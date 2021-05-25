package traffic

import "github.com/dtcookie/hcl"

// SpikeDetection The configuration of traffic spikes detection.
type SpikeDetection struct {
	Enabled             bool   `json:"enabled"`                       // The detection is enabled (`true`) or disabled (`false`).
	TrafficSpikePercent *int32 `json:"trafficSpikePercent,omitempty"` // Alert if the observed traffic is more than *X* % of the expected value.
}

func (me *SpikeDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "The detection is enabled (`true`) or disabled (`false`)",
		},
		"percent": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Alert if the observed traffic is less than *X* % of the expected value",
		},
	}
}
