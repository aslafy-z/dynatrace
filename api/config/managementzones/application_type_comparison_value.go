package managementzones

// ApplicationTypeComparisonValue The value to compare to.
type ApplicationTypeComparisonValue string

// ApplicationTypeComparisonValues offers the known enum values
var ApplicationTypeComparisonValues = struct {
	AgentlessMonitoring ApplicationTypeComparisonValue
	Amp                 ApplicationTypeComparisonValue
	AutoInjected        ApplicationTypeComparisonValue
	Default             ApplicationTypeComparisonValue
	SaasVendor          ApplicationTypeComparisonValue
}{
	"AGENTLESS_MONITORING",
	"AMP",
	"AUTO_INJECTED",
	"DEFAULT",
	"SAAS_VENDOR",
}
