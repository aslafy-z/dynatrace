package comparisoninfo

import "github.com/dtcookie/dynatrace/api/config/metrics/calculated/service/propagation"

// NumberRequestAttribute Comparison for `NUMBER_REQUEST_ATTRIBUTE` attributes, specifically of the generic **Number** type.
type NumberRequestAttribute struct {
	BaseComparisonInfo                                  // `json:",type=NUMBER_REQUEST_ATTRIBUTE"`
	MatchOnChildCalls  *bool                            `json:"matchOnChildCalls,omitempty"` // If `true`, the request attribute is matched on child service calls.    Default is `false`.
	RequestAttribute   string                           `json:"requestAttribute"`            // has no documentation
	Source             *propagation.Source              `json:"source,omitempty"`            // Defines valid sources of request attributes for conditions or placeholders.
	Value              *float64                         `json:"value,omitempty"`             // The value to compare to.
	Values             []float64                        `json:"values,omitempty"`            // The values to compare to.
	Comparison         NumberRequestAttributeComparison `json:"comparison"`                  // Operator of the comparision. You can reverse it by setting **negate** to `true`.
}

// NumberRequestAttributeComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type NumberRequestAttributeComparison string

// NumberRequestAttributeComparisons offers the known enum values
var NumberRequestAttributeComparisons = struct {
	Equals             NumberRequestAttributeComparison
	EqualsAnyOf        NumberRequestAttributeComparison
	Exists             NumberRequestAttributeComparison
	GreaterThan        NumberRequestAttributeComparison
	GreaterThanOrEqual NumberRequestAttributeComparison
	LowerThan          NumberRequestAttributeComparison
	LowerThanOrEqual   NumberRequestAttributeComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
	"GREATER_THAN",
	"GREATER_THAN_OR_EQUAL",
	"LOWER_THAN",
	"LOWER_THAN_OR_EQUAL",
}
