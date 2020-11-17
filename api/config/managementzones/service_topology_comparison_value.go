package managementzones

// ServiceTopologyComparisonValue The value to compare to.
type ServiceTopologyComparisonValue string

// ServiceTopologyComparisonValues offers the known enum values
var ServiceTopologyComparisonValues = struct {
	ExternalService ServiceTopologyComparisonValue
	FullyMonitored  ServiceTopologyComparisonValue
	OpaqueService   ServiceTopologyComparisonValue
}{
	"EXTERNAL_SERVICE",
	"FULLY_MONITORED",
	"OPAQUE_SERVICE",
}
