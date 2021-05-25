package traffic

import "github.com/dtcookie/hcl"

type Detection struct {
	Drops  *DropDetection  // The configuration of traffic drops detection.
	Spikes *SpikeDetection // The configuration of traffic spikes detection.
}

func (me *Detection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"drops": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The configuration of traffic drops detection",
			Elem:        &hcl.Resource{Schema: new(DropDetection).Schema()},
		},
		"spikes": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The configuration of traffic spikes detection",
			Elem:        &hcl.Resource{Schema: new(SpikeDetection).Schema()},
		},
	}
}
