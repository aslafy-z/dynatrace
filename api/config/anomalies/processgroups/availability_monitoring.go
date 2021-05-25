package processgroups

// AvailabilityMonitoringPG Configuration of the availability monitoring for the process group.
type AvailabilityMonitoring struct {
	Method           Method `json:"method"`                     // How to monitor the availability of the process group:  * `PROCESS_IMPACT`: Alert if any process of the group becomes unavailable.  * `MINIMUM_THRESHOLD`: Alert if the number of active processes in the group falls below the specified threshold.  * `OFF`: Availability monitoring is disabled.
	MinimumThreshold *int32 `json:"minimumThreshold,omitempty"` // Alert if the number of active processes in the group is lower than this value.
}

// Method How to monitor the availability of the process group:
// * `PROCESS_IMPACT`: Alert if any process of the group becomes unavailable.
// * `MINIMUM_THRESHOLD`: Alert if the number of active processes in the group falls below the specified threshold.
// * `OFF`: Availability monitoring is disabled.
type Method string

// Methods offers the known enum values
var Methods = struct {
	MinimumThreshold Method
	Off              Method
	ProcessImpact    Method
}{
	"MINIMUM_THRESHOLD",
	"OFF",
	"PROCESS_IMPACT",
}
