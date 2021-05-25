package databaseservices

// ConnectionFailureDetection Parameters of the failed database connections detection.
// The alert is triggered when failed connections number exceeds **connectionFailsCount** during any **timePeriodMinutes** minutes period.
type ConnectionFailureDetection struct {
	ConnectionFailsCount *int32 `json:"connectionFailsCount,omitempty"` // Number of failed database connections during any **timePeriodMinutes** minutes period to trigger an alert.
	Enabled              bool   `json:"enabled"`                        // The detection is enabled (`true`) or disabled (`false`).
	TimePeriodMinutes    *int32 `json:"timePeriodMinutes,omitempty"`    // The *X* minutes time period during which the **connectionFailsCount** is evaluated.
}
