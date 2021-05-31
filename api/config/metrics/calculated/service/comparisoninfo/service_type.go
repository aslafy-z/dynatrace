package comparisoninfo

// ServiceType Comparison for `SERVICE_TYPE` attributes.
type ServiceType struct {
	BaseComparisonInfo                       // `json:",type=SERVICE_TYPE"`
	Comparison         ServiceTypeComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *ServiceTypeValue     `json:"value,omitempty"`  // The value to compare to.
	Values             []ServiceTypeValue    `json:"values,omitempty"` // The values to compare to.
}

// ServiceTypeComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type ServiceTypeComparison string

// ServiceTypeComparisons offers the known enum values
var ServiceTypeComparisons = struct {
	Equals      ServiceTypeComparison
	EqualsAnyOf ServiceTypeComparison
	Exists      ServiceTypeComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// ServiceTypeValue The value to compare to.
type ServiceTypeValue string

// ServiceTypeValues offers the known enum values
var ServiceTypeValues = struct {
	BackgroundActivity          ServiceTypeValue
	CICSService                 ServiceTypeValue
	CustomService               ServiceTypeValue
	DatabaseService             ServiceTypeValue
	EnterpriseServiceBusService ServiceTypeValue
	External                    ServiceTypeValue
	IBMIntegrationBusService    ServiceTypeValue
	IMSService                  ServiceTypeValue
	MessagingService            ServiceTypeValue
	RMIService                  ServiceTypeValue
	RPCService                  ServiceTypeValue
	WebRequestService           ServiceTypeValue
	WebService                  ServiceTypeValue
}{
	"BACKGROUND_ACTIVITY",
	"CICS_SERVICE",
	"CUSTOM_SERVICE",
	"DATABASE_SERVICE",
	"ENTERPRISE_SERVICE_BUS_SERVICE",
	"EXTERNAL",
	"IBM_INTEGRATION_BUS_SERVICE",
	"IMS_SERVICE",
	"MESSAGING_SERVICE",
	"RMI_SERVICE",
	"RPC_SERVICE",
	"WEB_REQUEST_SERVICE",
	"WEB_SERVICE",
}
