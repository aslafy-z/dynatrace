package services

// LoadThreshold Minimal service load to detect response time degradation.
//  Response time degradation of services with smaller load won't trigger alerts.
type LoadThreshold string

// LoadThresholds offers the known enum values
var LoadThresholds = struct {
	FifteenRequestsPerMinute LoadThreshold
	FiveRequestsPerMinute    LoadThreshold
	OneRequestPerMinute      LoadThreshold
	TenRequestsPerMinute     LoadThreshold
}{
	"FIFTEEN_REQUESTS_PER_MINUTE",
	"FIVE_REQUESTS_PER_MINUTE",
	"ONE_REQUEST_PER_MINUTE",
	"TEN_REQUESTS_PER_MINUTE",
}
