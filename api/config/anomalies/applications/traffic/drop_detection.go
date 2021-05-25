package traffic

import "github.com/dtcookie/hcl"

// DropDetection The configuration of traffic drops detection.
type DropDetection struct {
	Enabled            bool   `json:"enabled"`                      // The detection is enabled (`true`) or disabled (`false`).
	TrafficDropPercent *int32 `json:"trafficDropPercent,omitempty"` // Alert if the observed traffic is less than *X* % of the expected value.
}

func (me *DropDetection) Schema() map[string]*hcl.Schema {
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
