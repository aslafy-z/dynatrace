package managementzones

// DatabaseTopologyComparisonValue The value to compare to.
type DatabaseTopologyComparisonValue string

// DatabaseTopologyComparisonValues offers the known enum values
var DatabaseTopologyComparisonValues = struct {
	Cluster       DatabaseTopologyComparisonValue
	Embedded      DatabaseTopologyComparisonValue
	Failover      DatabaseTopologyComparisonValue
	Ipc           DatabaseTopologyComparisonValue
	LoadBalancing DatabaseTopologyComparisonValue
	SingleServer  DatabaseTopologyComparisonValue
	Unspecified   DatabaseTopologyComparisonValue
}{
	"CLUSTER",
	"EMBEDDED",
	"FAILOVER",
	"IPC",
	"LOAD_BALANCING",
	"SINGLE_SERVER",
	"UNSPECIFIED",
}
