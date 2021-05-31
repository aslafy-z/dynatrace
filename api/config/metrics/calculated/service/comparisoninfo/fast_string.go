package comparisoninfo

// FastString Comparison for `FAST_STRING` attributes. Use it for all service property attributes.
type FastString struct {
	BaseComparisonInfo                      // `json:",type=FAST_STRING"`
	CaseSensitive      *bool                `json:"caseSensitive,omitempty"` // The comparison is case-sensitive (`true`) or not case-sensitive (`false`).
	Comparison         FastStringComparison `json:"comparison"`              // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *string              `json:"value,omitempty"`         // The value to compare to.
	Values             []string             `json:"values,omitempty"`        // The values to compare to.
}

// FastStringComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type FastStringComparison string

// FastStringComparisons offers the known enum values
var FastStringComparisons = struct {
	Contains    FastStringComparison
	Equals      FastStringComparison
	EqualsAnyOf FastStringComparison
}{
	"CONTAINS",
	"EQUALS",
	"EQUALS_ANY_OF",
}
