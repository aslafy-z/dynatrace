package autotags

// AzureSkuComparisionValue The value to compare to.
type AzureSkuComparisionValue string

// AzureSkuComparisionValues offers the known enum values
var AzureSkuComparisionValues = struct {
	Basic    AzureSkuComparisionValue
	Dynamic  AzureSkuComparisionValue
	Free     AzureSkuComparisionValue
	Premium  AzureSkuComparisionValue
	Shared   AzureSkuComparisionValue
	Standard AzureSkuComparisionValue
}{
	"BASIC",
	"DYNAMIC",
	"FREE",
	"PREMIUM",
	"SHARED",
	"STANDARD",
}
