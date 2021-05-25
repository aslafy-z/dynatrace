package processgroups

// AnomalyDetection Configuration of anomaly detection for the process group.
type AnomalyDetection struct {
	AvailabilityMonitoring *AvailabilityMonitoring `json:"availabilityMonitoring,omitempty"` // Configuration of the availability monitoring for the process group.
}
