package alertingprofiles

// Operator Operator of the comparison.
//  You can reverse it by setting **negate** to `true`.
type Operator string

// Operators offers the known enum values
var Operators = struct {
	BeginsWith    Operator
	Contains      Operator
	ContainsRegex Operator
	EndsWith      Operator
	Equals        Operator
}{
	"BEGINS_WITH",
	"CONTAINS",
	"CONTAINS_REGEX",
	"ENDS_WITH",
	"EQUALS",
}
