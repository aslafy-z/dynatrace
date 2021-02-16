package autotags

// AzureComputeModeComparisonValue The value to compare to.
type AzureComputeModeComparisonValue string

// AzureComputeModeComparisonValues offers the known enum values
var AzureComputeModeComparisonValues = struct {
	Dedicated AzureComputeModeComparisonValue
	Shared    AzureComputeModeComparisonValue
}{
	"DEDICATED",
	"SHARED",
}
