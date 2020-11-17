package managementzones

// OSTypeComparisonOperator Operator of the comparison. You can reverse it by setting **negate** to `true`.
// Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
type OSTypeComparisonOperator string

// OSTypeComparisonOperators offers the known enum values
var OSTypeComparisonOperators = struct {
	Equals OSTypeComparisonOperator
	Exists OSTypeComparisonOperator
}{
	"EQUALS",
	"EXISTS",
}
