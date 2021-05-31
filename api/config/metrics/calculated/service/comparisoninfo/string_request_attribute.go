package comparisoninfo

import "github.com/dtcookie/dynatrace/api/config/metrics/calculated/service/propagation"

// StringRequestAttribute Comparison for `STRING_REQUEST_ATTRIBUTE` attributes, specifically of the **String** type.
type StringRequestAttribute struct {
	BaseComparisonInfo                                  // `json:",type=STRING_REQUEST_ATTRIBUTE"`
	Source             *propagation.Source              `json:"source,omitempty"`            // Defines valid sources of request attributes for conditions or placeholders.
	Value              *string                          `json:"value,omitempty"`             // The value to compare to.
	Values             []string                         `json:"values,omitempty"`            // The values to compare to.
	CaseSensitive      *bool                            `json:"caseSensitive,omitempty"`     // The comparison is case-sensitive (`true`) or not case-sensitive (`false`).
	Comparison         StringRequestAttributeComparison `json:"comparison"`                  // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	MatchOnChildCalls  *bool                            `json:"matchOnChildCalls,omitempty"` // If `true`, the request attribute is matched on child service calls.   Default is `false`.
	RequestAttribute   string                           `json:"requestAttribute"`            // has no documentation
}

// StringRequestAttributeComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type StringRequestAttributeComparison string

// StringRequestAttributeComparisons offers the known enum values
var StringRequestAttributeComparisons = struct {
	BeginsWith      StringRequestAttributeComparison
	BeginsWithAnyOf StringRequestAttributeComparison
	Contains        StringRequestAttributeComparison
	EndsWith        StringRequestAttributeComparison
	EndsWithAnyOf   StringRequestAttributeComparison
	Equals          StringRequestAttributeComparison
	EqualsAnyOf     StringRequestAttributeComparison
	Exists          StringRequestAttributeComparison
	RegexMatches    StringRequestAttributeComparison
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
