package managementzones

// OSArchitectureComparisonOperator Operator of the comparison. You can reverse it by setting **negate** to `true`.
// Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
type OSArchitectureComparisonOperator string

// OSArchitectureComparisonOperators offers the known enum values
var OSArchitectureComparisonOperators = struct {
	Equals OSArchitectureComparisonOperator
	Exists OSArchitectureComparisonOperator
}{
	"EQUALS",
	"EXISTS",
}
