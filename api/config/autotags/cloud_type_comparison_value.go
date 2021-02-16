package autotags

// CloudTypeComparisonValue The value to compare to.
type CloudTypeComparisonValue string

// CloudTypeComparisonValues offers the known enum values
var CloudTypeComparisonValues = struct {
	Azure               CloudTypeComparisonValue
	EC2                 CloudTypeComparisonValue
	GoogleCloudPlatform CloudTypeComparisonValue
	OpenStack           CloudTypeComparisonValue
	Oracle              CloudTypeComparisonValue
	Unrecognized        CloudTypeComparisonValue
}{
	"AZURE",
	"EC2",
	"GOOGLE_CLOUD_PLATFORM",
	"OPENSTACK",
	"ORACLE",
	"UNRECOGNIZED",
}
