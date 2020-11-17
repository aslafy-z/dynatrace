package managementzones

// OSTypeComparisonValue The value to compare to.
type OSTypeComparisonValue string

// OSTypeComparisonValues offers the known enum values
var OSTypeComparisonValues = struct {
	AIX     OSTypeComparisonValue
	Darwin  OSTypeComparisonValue
	Hpux    OSTypeComparisonValue
	Linux   OSTypeComparisonValue
	Solaris OSTypeComparisonValue
	Windows OSTypeComparisonValue
	Zos     OSTypeComparisonValue
}{
	"AIX",
	"DARWIN",
	"HPUX",
	"LINUX",
	"SOLARIS",
	"WINDOWS",
	"ZOS",
}
