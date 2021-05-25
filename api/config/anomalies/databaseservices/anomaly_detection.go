package databaseservices

import (
	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/failurerate"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/load"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/responsetime"
)

// AnomalyDetection The configuration of the anomaly detection for database services.
type AnomalyDetection struct {
	FailureRateIncrease            *failurerate.Detection      `json:"failureRateIncrease"`            // Configuration of failure rate increase detection.
	LoadDrop                       *load.DropDetection         `json:"loadDrop,omitempty"`             // The configuration of load drops detection.
	LoadSpike                      *load.SpikeDetection        `json:"loadSpike,omitempty"`            // The configuration of load spikes detection.
	ResponseTimeDegradation        *responsetime.Detection     `json:"responseTimeDegradation"`        // Configuration of response time degradation detection.
	DatabaseConnectionFailureCount *ConnectionFailureDetection `json:"databaseConnectionFailureCount"` // Parameters of the failed database connections detection.  The alert is triggered when failed connections number exceeds **connectionFailsCount** during any **timePeriodMinutes** minutes period.
	Metadata                       *api.ConfigMetadata         `json:"metadata,omitempty"`             // Metadata useful for debugging
}
