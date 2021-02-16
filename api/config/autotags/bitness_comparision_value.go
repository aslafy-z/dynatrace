package autotags

// BitnessComparisionValue The value to compare to.
type BitnessComparisionValue string

// BitnessComparisionValues offers the known enum values
var BitnessComparisionValues = struct {
	V32 BitnessComparisionValue
	V64 BitnessComparisionValue
}{
	"32",
	"64",
}
