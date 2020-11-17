package requestattributes

// Operator Operator comparing the extracted value to the comparison value.
type Operator string

// Operators offers the known enum values
var Operators = struct {
	BeginsWith Operator
	Contains   Operator
	EndsWith   Operator
	Equals     Operator
}{
	"BEGINS_WITH",
	"CONTAINS",
	"ENDS_WITH",
	"EQUALS",
}
