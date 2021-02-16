package autotags

// ServiceTypeComparisonOperator Operator of the comparison. You can reverse it by setting **negate** to `true`.
// Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
type ServiceTypeComparisonOperator string

// ServiceTypeComparisonOperators offers the known enum values
var ServiceTypeComparisonOperators = struct {
	Equals ServiceTypeComparisonOperator
	Exists ServiceTypeComparisonOperator
}{
	"EQUALS",
	"EXISTS",
}
