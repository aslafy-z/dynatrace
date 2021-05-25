package applications

import (
	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/api/config/anomalies/applications/traffic"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/failurerate"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/responsetime"
	"github.com/dtcookie/hcl"
)

// AnomalyDetection The configuration of anomaly detection for applications.
type AnomalyDetection struct {
	ResponseTimeDegradation *responsetime.Detection `json:"responseTimeDegradation"` // Configuration of response time degradation detection.
	TrafficDrop             *traffic.DropDetection  `json:"trafficDrop"`             // The configuration of traffic drops detection.
	TrafficSpike            *traffic.SpikeDetection `json:"trafficSpike"`            // The configuration of traffic spikes detection.
	FailureRateIncrease     *failurerate.Detection  `json:"failureRateIncrease"`     // Configuration of failure rate increase detection.
	Metadata                *api.ConfigMetadata     `json:"metadata,omitempty"`      // Metadata useful for debugging
}

func (me *AnomalyDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"traffic": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for anomalies regarding traffic",
			Elem:        &hcl.Resource{Schema: new(traffic.Detection).Schema()},
		},
		"response_time": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of response time degradation detection",
			Elem:        &hcl.Resource{Schema: new(responsetime.Detection).Schema()},
		},
		"failure_rate": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of failure rate increase detection",
			Elem:        &hcl.Resource{Schema: new(failurerate.Detection).Schema()},
		},
	}
}
