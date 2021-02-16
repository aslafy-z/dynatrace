package autotags

// StringComparisonOperator Operator of the comparison. You can reverse it by setting **negate** to `true`.
// Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
type StringComparisonOperator string

// StringComparisonOperators offers the known enum values
var StringComparisonOperators = struct {
	BeginsWith   StringComparisonOperator
	Contains     StringComparisonOperator
	EndsWith     StringComparisonOperator
	Equals       StringComparisonOperator
	Exists       StringComparisonOperator
	RegexMatches StringComparisonOperator
}{
	"BEGINS_WITH",
	"CONTAINS",
	"ENDS_WITH",
	"EQUALS",
	"EXISTS",
	"REGEX_MATCHES",
}
