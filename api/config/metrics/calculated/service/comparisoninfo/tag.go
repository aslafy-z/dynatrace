package comparisoninfo

// Tag Comparison for `TAG` attributes.
type Tag struct {
	BaseComparisonInfo               // `json:",type=TAG"`
	Comparison         TagComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *TagInfo      `json:"value,omitempty"`  // Tag of a Dynatrace entity.
	Values             []*TagInfo    `json:"values,omitempty"` // The values to compare to.
}

// TagComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type TagComparison string

// TagComparisons offers the known enum values
var TagComparisons = struct {
	Equals            TagComparison
	EqualsAnyOf       TagComparison
	TagKeyEquals      TagComparison
	TagKeyEqualsAnyOf TagComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"TAG_KEY_EQUALS",
	"TAG_KEY_EQUALS_ANY_OF",
}
