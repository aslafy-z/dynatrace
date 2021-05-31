package comparisoninfo

// Number Comparison for `NUMBER` attributes.
type Number struct {
	BaseComparisonInfo                  // `json:",type=NUMBER"`
	Comparison         NumberComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *float64         `json:"value,omitempty"`  // The value to compare to.
	Values             []float64        `json:"values,omitempty"` // The values to compare to.
}

// NumberComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type NumberComparison string

// NumberComparisons offers the known enum values
var NumberComparisons = struct {
	Equals             NumberComparison
	EqualsAnyOf        NumberComparison
	Exists             NumberComparison
	GreaterThan        NumberComparison
	GreaterThanOrEqual NumberComparison
	LowerThan          NumberComparison
	LowerThanOrEqual   NumberComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
	"GREATER_THAN",
	"GREATER_THAN_OR_EQUAL",
	"LOWER_THAN",
	"LOWER_THAN_OR_EQUAL",
}
