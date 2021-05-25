package diskevents

import (
	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/api/config/common"
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
	TagFilters       []*common.TagFilter `json:"tagFilters,omitempty"`     // Narrows the rule usage down to the hosts matching the specified tags.
	Metadata         *api.ConfigMetadata `json:"metadata,omitempty"`       // Metadata useful for debugging
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
