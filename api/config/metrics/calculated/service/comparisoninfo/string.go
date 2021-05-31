package comparisoninfo

// String Comparison for `STRING` attributes.
type String struct {
	BaseComparisonInfo                  // `json:",type=STRING"`
	CaseSensitive      *bool            `json:"caseSensitive,omitempty"` // The comparison is case-sensitive (`true`) or not case-sensitive (`false`).
	Comparison         StringComparison `json:"comparison"`              // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *string          `json:"value,omitempty"`         // The value to compare to.
	Values             []string         `json:"values,omitempty"`        // The values to compare to.
}

// StringComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type StringComparison string

// StringComparisons offers the known enum values
var StringComparisons = struct {
	BeginsWith      StringComparison
	BeginsWithAnyOf StringComparison
	Contains        StringComparison
	EndsWith        StringComparison
	EndsWithAnyOf   StringComparison
	Equals          StringComparison
	EqualsAnyOf     StringComparison
	Exists          StringComparison
	RegexMatches    StringComparison
}{
	"BEGINS_WITH",
	"BEGINS_WITH_ANY_OF",
	"CONTAINS",
	"ENDS_WITH",
	"ENDS_WITH_ANY_OF",
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
	"REGEX_MATCHES",
}
