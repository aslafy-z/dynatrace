package comparisoninfo

// FlawState Comparison for `FLAW_STATE` attributes.
type FlawState struct {
	BaseComparisonInfo                     // `json:",type=FLAW_STATE"`
	Comparison         FlawStateComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *FlawStateValue     `json:"value,omitempty"`  // The value to compare to.
	Values             []FlawStateValue    `json:"values,omitempty"` // The values to compare to.
}

// FlawStateComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type FlawStateComparison string

// FlawStateComparisons offers the known enum values
var FlawStateComparisons = struct {
	Equals      FlawStateComparison
	EqualsAnyOf FlawStateComparison
	Exists      FlawStateComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// FlawStateValue The value to compare to.
type FlawStateValue string

// FlawStateValues offers the known enum values
var FlawStateValues = struct {
	Flawed    FlawStateValue
	NotFlawed FlawStateValue
}{
	"FLAWED",
	"NOT_FLAWED",
}
