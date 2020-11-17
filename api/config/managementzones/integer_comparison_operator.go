package managementzones

// IntegerComparisonOperator Operator of the comparison. You can reverse it by setting **negate** to `true`.
// Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
type IntegerComparisonOperator string

// IntegerComparisonOperators offers the known enum values
var IntegerComparisonOperators = struct {
	Equals             IntegerComparisonOperator
	Exists             IntegerComparisonOperator
	GreaterThan        IntegerComparisonOperator
	GreaterThanOrEqual IntegerComparisonOperator
	LowerThan          IntegerComparisonOperator
	LowerThanOrEqual   IntegerComparisonOperator
}{
	"EQUALS",
	"EXISTS",
	"GREATER_THAN",
	"GREATER_THAN_OR_EQUAL",
	"LOWER_THAN",
	"LOWER_THAN_OR_EQUAL",
}
