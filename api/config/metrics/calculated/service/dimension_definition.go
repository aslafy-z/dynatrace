package service

// DimensionDefinition Parameters of a definition of a calculated service metric.
type DimensionDefinition struct {
	TopXDirection   TopXDirection   `json:"topXDirection"`          // How to calculate the **topX** values.
	Dimension       string          `json:"dimension"`              // The dimension value pattern.   You can define custom placeholders in the **placeholders** field and use them here.
	Name            string          `json:"name"`                   // The name of the dimension.
	Placeholders    []*Placeholder  `json:"placeholders,omitempty"` // The list of custom placeholders to be used in a dimension value pattern.
	TopX            int32           `json:"topX"`                   // The number of top values to be calculated.
	TopXAggregation TopXAggregation `json:"topXAggregation"`        // The aggregation of the dimension.
}

// TopXDirection How to calculate the **topX** values.
type TopXDirection string

// TopXDirections offers the known enum values
var TopXDirections = struct {
	Ascending  TopXDirection
	Descending TopXDirection
}{
	"ASCENDING",
	"DESCENDING",
}

// TopXAggregation The aggregation of the dimension.
type TopXAggregation string

// TopXAggregations offers the known enum values
var TopXAggregations = struct {
	Average         TopXAggregation
	Count           TopXAggregation
	Max             TopXAggregation
	Min             TopXAggregation
	OfInterestRatio TopXAggregation
	OtherRatio      TopXAggregation
	SingleValue     TopXAggregation
	Sum             TopXAggregation
}{
	"AVERAGE",
	"COUNT",
	"MAX",
	"MIN",
	"OF_INTEREST_RATIO",
	"OTHER_RATIO",
	"SINGLE_VALUE",
	"SUM",
}
