package autotags

// OSArchitectureComparisonValue The value to compare to.
type OSArchitectureComparisonValue string

// OSArchitectureComparisonValues offers the known enum values
var OSArchitectureComparisonValues = struct {
	Arm    OSArchitectureComparisonValue
	Ia64   OSArchitectureComparisonValue
	Parisc OSArchitectureComparisonValue
	Ppc    OSArchitectureComparisonValue
	Ppcle  OSArchitectureComparisonValue
	S390   OSArchitectureComparisonValue
	Sparc  OSArchitectureComparisonValue
	X86    OSArchitectureComparisonValue
	Zos    OSArchitectureComparisonValue
}{
	"ARM",
	"IA64",
	"PARISC",
	"PPC",
	"PPCLE",
	"S390",
	"SPARC",
	"X86",
	"ZOS",
}
