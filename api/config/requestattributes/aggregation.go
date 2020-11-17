package requestattributes

// Aggregation Aggregation type for the request values.
type Aggregation string

// Aggregations offers the known enum values
var Aggregations = struct {
	AllDistinctValues   Aggregation
	Average             Aggregation
	CountDistinctValues Aggregation
	CountValues         Aggregation
	First               Aggregation
	Last                Aggregation
	Maximum             Aggregation
	Minimum             Aggregation
	Sum                 Aggregation
}{
	"ALL_DISTINCT_VALUES",
	"AVERAGE",
	"COUNT_DISTINCT_VALUES",
	"COUNT_VALUES",
	"FIRST",
	"LAST",
	"MAXIMUM",
	"MINIMUM",
	"SUM",
}
