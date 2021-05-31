package service

import (
	"github.com/dtcookie/hcl"
)

// CalculatedMetricDefinition The definition of a calculated service metric.
type CalculatedMetricDefinition struct {
	Metric           Metric  `json:"metric"`                     // The metric to be captured.
	RequestAttribute *string `json:"requestAttribute,omitempty"` // The request attribute to be captured. Only applicable when the **metric** parameter is set to `REQUEST_ATTRIBUTE`.
}

func (me *CalculatedMetricDefinition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"metric": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The metric to be captured. Possible values are `CPU_TIME`, `DATABASE_CHILD_CALL_COUNT`, `DATABASE_CHILD_CALL_TIME`, `DISK_IO_TIME`, `EXCEPTION_COUNT`, `FAILED_REQUEST_COUNT`, `FAILED_REQUEST_COUNT_CLIENT`, `FAILURE_RATE`, `FAILURE_RATE_CLIENT`, `HTTP_4XX_ERROR_COUNT`, `HTTP_4XX_ERROR_COUNT_CLIENT`, `HTTP_5XX_ERROR_COUNT`, `HTTP_5XX_ERROR_COUNT_CLIENT`, `IO_TIME`, `LOCK_TIME`, `NETWORK_IO_TIME`, `NON_DATABASE_CHILD_CALL_COUNT`, `NON_DATABASE_CHILD_CALL_TIME`, `PROCESSING_TIME`, `REQUEST_ATTRIBUTE`, `REQUEST_COUNT`, `RESPONSE_TIME`, `RESPONSE_TIME_CLIENT`, `SUCCESSFUL_REQUEST_COUNT`, `SUCCESSFUL_REQUEST_COUNT_CLIENT` and `WAIT_TIME`",
		},
		"request_attribute": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The request attribute to be captured. Only applicable when the **metric** parameter is set to `REQUEST_ATTRIBUTE`",
		},
	}
}

// Metric The metric to be captured.
type Metric string

// Metrics offers the known enum values
var Metrics = struct {
	CPUTime                      Metric
	DatabaseChildCallCount       Metric
	DatabaseChildCallTime        Metric
	DiskIoTime                   Metric
	ExceptionCount               Metric
	FailedRequestCount           Metric
	FailedRequestCountClient     Metric
	FailureRate                  Metric
	FailureRateClient            Metric
	Http4xxErrorCount            Metric
	Http4xxErrorCountClient      Metric
	Http5xxErrorCount            Metric
	Http5xxErrorCountClient      Metric
	IoTime                       Metric
	LockTime                     Metric
	NetworkIoTime                Metric
	NonDatabaseChildCallCount    Metric
	NonDatabaseChildCallTime     Metric
	ProcessingTime               Metric
	RequestAttribute             Metric
	RequestCount                 Metric
	ResponseTime                 Metric
	ResponseTimeClient           Metric
	SuccessfulRequestCount       Metric
	SuccessfulRequestCountClient Metric
	WaitTime                     Metric
}{
	"CPU_TIME",
	"DATABASE_CHILD_CALL_COUNT",
	"DATABASE_CHILD_CALL_TIME",
	"DISK_IO_TIME",
	"EXCEPTION_COUNT",
	"FAILED_REQUEST_COUNT",
	"FAILED_REQUEST_COUNT_CLIENT",
	"FAILURE_RATE",
	"FAILURE_RATE_CLIENT",
	"HTTP_4XX_ERROR_COUNT",
	"HTTP_4XX_ERROR_COUNT_CLIENT",
	"HTTP_5XX_ERROR_COUNT",
	"HTTP_5XX_ERROR_COUNT_CLIENT",
	"IO_TIME",
	"LOCK_TIME",
	"NETWORK_IO_TIME",
	"NON_DATABASE_CHILD_CALL_COUNT",
	"NON_DATABASE_CHILD_CALL_TIME",
	"PROCESSING_TIME",
	"REQUEST_ATTRIBUTE",
	"REQUEST_COUNT",
	"RESPONSE_TIME",
	"RESPONSE_TIME_CLIENT",
	"SUCCESSFUL_REQUEST_COUNT",
	"SUCCESSFUL_REQUEST_COUNT_CLIENT",
	"WAIT_TIME",
}
