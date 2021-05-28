package diskevents

import (
	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/api/config/common"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// AnomalyDetection has no documentation
type AnomalyDetection struct {
	ID               *string             `json:"id,omitempty"`             // The ID of the disk event rule.
	Name             string              `json:"name"`                     // The name of the disk event rule.
	HostGroupID      *string             `json:"hostGroupId,omitempty"`    // Narrows the rule usage down to disks that run on hosts that themselves run on the specified host group.
	Threshold        float64             `json:"threshold"`                // The threshold to trigger disk event.   * A percentage for `LowDiskSpace` or `LowInodes` metrics.   * In milliseconds for `ReadTimeExceeding` or `WriteTimeExceeding` metrics.
	DiskNameFilter   *DiskNameFilter     `json:"diskNameFilter,omitempty"` // Narrows the rule usage down to disks, matching the specified criteria.
	Enabled          bool                `json:"enabled"`                  // Disk event rule enabled/disabled.
	Samples          int32               `json:"samples"`                  // The number of samples to evaluate.
	ViolatingSamples int32               `json:"violatingSamples"`         // The number of samples that must violate the threshold to trigger an event. Must not exceed the number of evaluated samples.
	Metric           Metric              `json:"metric"`                   // The metric to monitor.
	TagFilters       common.TagFilters   `json:"tagFilters,omitempty"`     // Narrows the rule usage down to the hosts matching the specified tags.
	Metadata         *api.ConfigMetadata `json:"metadata,omitempty"`       // Metadata useful for debugging
}

func (me *AnomalyDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the disk event rule",
		},
		"host_group_id": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "Narrows the rule usage down to disks that run on hosts that themselves run on the specified host group",
		},
		"threshold": {
			Type:        hcl.TypeFloat,
			Required:    true,
			Description: "The threshold to trigger disk event.   * A percentage for `LowDiskSpace` or `LowInodes` metrics.   * In milliseconds for `ReadTimeExceeding` or `WriteTimeExceeding` metrics",
		},
		"disk_name": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Narrows the rule usage down to disks, matching the specified criteria",
			Elem:        &hcl.Resource{Schema: new(DiskNameFilter).Schema()},
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "Disk event rule enabled/disabled",
		},
		"samples": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "The number of samples to evaluate",
		},
		"violating_samples": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "The number of samples that must violate the threshold to trigger an event. Must not exceed the number of evaluated samples",
		},
		"metric": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The metric to monitor",
		},
		"tags": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Narrows the rule usage down to the hosts matching the specified tags",
			Elem:        &hcl.Resource{Schema: new(common.TagFilters).Schema()},
		},
	}
}

func (me *AnomalyDetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	properties, err := decoder.MarshalAll(map[string]interface{}{
		"name":              me.Name,
		"host_group_id":     me.HostGroupID,
		"threshold":         me.Threshold,
		"enabled":           me.Enabled,
		"violating_samples": me.ViolatingSamples,
		"metric":            me.Metric,
		"tags":              me.TagFilters,
	})
	if err != nil {
		return nil, err
	}
	if me.DiskNameFilter != nil {
		if marshalled, err := me.DiskNameFilter.MarshalHCL(hcl.NewDecoder(decoder, "disk_name", 0)); err != nil {
			return nil, err
		} else {
			properties["disk_name"] = []interface{}{marshalled}
		}

	}
	return properties, nil
}

func (me *AnomalyDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	reader := hcl.NewReader(decoder, nil)
	me.Name = opt.String(reader.String("name"))
	me.HostGroupID = reader.String("host_group_id")
	me.Threshold = opt.Float64(reader.Float64("threshold"))
	me.Enabled = opt.Bool(reader.Bool("enabled"))
	me.ViolatingSamples = opt.Int32(reader.Int32("violating_samples"))
	me.Metric = Metric(opt.String(reader.String("metric")))
	if _, ok := decoder.GetOk("disk_name.#"); ok {
		me.DiskNameFilter = new(DiskNameFilter)
		if err := me.DiskNameFilter.UnmarshalHCL(hcl.NewDecoder(decoder, "disk_name", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("tags.#"); ok {
		me.TagFilters = common.TagFilters{}
		if err := me.TagFilters.UnmarshalHCL(hcl.NewDecoder(decoder, "tags", 0)); err != nil {
			return err
		}
	}
	return nil
}

// Metric The metric to monitor.
type Metric string

// Metrics offers the known enum values
var Metrics = struct {
	LowDiskSpace       Metric
	LowInodes          Metric
	ReadTimeExceeding  Metric
	WriteTimeExceeding Metric
}{
	"LOW_DISK_SPACE",
	"LOW_INODES",
	"READ_TIME_EXCEEDING",
	"WRITE_TIME_EXCEEDING",
}
