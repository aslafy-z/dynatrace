package comparisoninfo

// FailureReason Comparison for `FAILURE_REASON` attributes.
type FailureReason struct {
	BaseComparisonInfo                         // `json:",type=FAILURE_REASON"`
	Comparison         FailureReasonComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *FailureReasonValue     `json:"value,omitempty"`  // The value to compare to.
	Values             []FailureReasonValue    `json:"values,omitempty"` // The values to compare to.
}

// FailureReasonComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type FailureReasonComparison string

// FailureReasonComparisons offers the known enum values
var FailureReasonComparisons = struct {
	Equals      FailureReasonComparison
	EqualsAnyOf FailureReasonComparison
	Exists      FailureReasonComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// FailureReasonValue The value to compare to.
type FailureReasonValue string

// FailureReasonValues offers the known enum values
var FailureReasonValues = struct {
	ExceptionAtEntryNode FailureReasonValue
	ExceptionOnAnyNode   FailureReasonValue
	HTTPCode             FailureReasonValue
	RequestAttribute     FailureReasonValue
}{
	"EXCEPTION_AT_ENTRY_NODE",
	"EXCEPTION_ON_ANY_NODE",
	"HTTP_CODE",
	"REQUEST_ATTRIBUTE",
}
