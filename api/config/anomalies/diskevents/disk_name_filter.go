package diskevents

// DiskNameFilter Narrows the rule usage down to disks, matching the specified criteria.
type DiskNameFilter struct {
	Operator Operator `json:"operator"` // Comparison operator.
	Value    string   `json:"value"`    // Value to compare to.
}

// Operator Comparison operator.
type Operator string

// Operators offers the known enum values
var Operators = struct {
	Contains         Operator
	DoesNotContain   Operator
	DoesNotEqual     Operator
	DoesNotStartWith Operator
	Equals           Operator
	StartsWith       Operator
}{
	"CONTAINS",
	"DOES_NOT_CONTAIN",
	"DOES_NOT_EQUAL",
	"DOES_NOT_START_WITH",
	"EQUALS",
	"STARTS_WITH",
}
