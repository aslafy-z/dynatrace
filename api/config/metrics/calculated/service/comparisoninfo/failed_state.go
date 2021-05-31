package comparisoninfo

// FailedState Comparison for `FAILED_STATE` attributes.
type FailedState struct {
	BaseComparisonInfo                       // `json:",type=FAILED_STATE"`
	Comparison         FailedStateComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *FailedStateValue     `json:"value,omitempty"`  // The value to compare to.
	Values             []FailedStateValue    `json:"values,omitempty"` // The values to compare to.
}

func (me *FailedState) GetType() Type {
	return Types.FailedState
}

// FailedStateValue The value to compare to.
type FailedStateValue string

// FailedStateValues offers the known enum values
var FailedStateValues = struct {
	Failed     FailedStateValue
	Successful FailedStateValue
}{
	"FAILED",
	"SUCCESSFUL",
}

// FailedStateComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type FailedStateComparison string

// FailedStateComparisons offers the known enum values
var FailedStateComparisons = struct {
	Equals      FailedStateComparison
	EqualsAnyOf FailedStateComparison
	Exists      FailedStateComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}
