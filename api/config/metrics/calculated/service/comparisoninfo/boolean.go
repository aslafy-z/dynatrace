package comparisoninfo

// Boolean Comparison for `BOOLEAN` attributes.
type Boolean struct {
	BaseComparisonInfo                   // `json:",type=BOOLEAN"`
	Values             []bool            `json:"values,omitempty"` // The values to compare to.
	Comparison         BooleanComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *bool             `json:"value,omitempty"`  // The value to compare to.
}

func (me *Boolean) GetType() Type {
	return Types.Boolean
}

// BooleanComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type BooleanComparison string

// BooleanComparisons offers the known enum values
var BooleanComparisons = struct {
	Equals      BooleanComparison
	EqualsAnyOf BooleanComparison
	Exists      BooleanComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}
