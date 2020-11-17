package managementzones

// ServiceTypeComparisonValue The value to compare to.
type ServiceTypeComparisonValue string

// ServiceTypeComparisonValues offers the known enum values
var ServiceTypeComparisonValues = struct {
	BackgroundActivity          ServiceTypeComparisonValue
	CICSService                 ServiceTypeComparisonValue
	CustomService               ServiceTypeComparisonValue
	DatabaseService             ServiceTypeComparisonValue
	EnterpriseServiceBusService ServiceTypeComparisonValue
	External                    ServiceTypeComparisonValue
	IBMIntegrationBusService    ServiceTypeComparisonValue
	IMSService                  ServiceTypeComparisonValue
	MessagingService            ServiceTypeComparisonValue
	QueueListenerService        ServiceTypeComparisonValue
	RMIService                  ServiceTypeComparisonValue
	RPCService                  ServiceTypeComparisonValue
	WebRequestService           ServiceTypeComparisonValue
	WebService                  ServiceTypeComparisonValue
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
	"QUEUE_LISTENER_SERVICE",
	"RMI_SERVICE",
	"RPC_SERVICE",
	"WEB_REQUEST_SERVICE",
	"WEB_SERVICE",
}
