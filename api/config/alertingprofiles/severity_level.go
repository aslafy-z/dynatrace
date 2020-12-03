package alertingprofiles

// SeverityLevel The severity level to trigger the alert.
type SeverityLevel string

// SeverityLevels offers the known enum values
var SeverityLevels = struct {
	Availability          SeverityLevel
	CustomAlert           SeverityLevel
	Error                 SeverityLevel
	MonitoringUnavailable SeverityLevel
	Performance           SeverityLevel
	ResourceContention    SeverityLevel
}{
	"AVAILABILITY",
	"CUSTOM_ALERT",
	"ERROR",
	"MONITORING_UNAVAILABLE",
	"PERFORMANCE",
	"RESOURCE_CONTENTION",
}
