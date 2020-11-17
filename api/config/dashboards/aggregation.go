package dashboards

import "encoding/json"

// Aggregation has no documentation
type Aggregation string

// Aggregations offers the known enum values
var Aggregations = struct {
	Avg              Aggregation
	Count            Aggregation
	Distinct         Aggregation
	Fastest10percent Aggregation
	Max              Aggregation
	Median           Aggregation
	Min              Aggregation
	None             Aggregation
	OfInterestRatio  Aggregation
	OtherRatio       Aggregation
	Percentile       Aggregation
	PerMin           Aggregation
	Slowest10percent Aggregation
	Slowest5percent  Aggregation
	Sum              Aggregation
	SumDimensions    Aggregation
}{
	"AVG",
	"COUNT",
	"DISTINCT",
	"FASTEST10PERCENT",
	"MAX",
	"MEDIAN",
	"MIN",
	"NONE",
	"OF_INTEREST_RATIO",
	"OTHER_RATIO",
	"PERCENTILE",
	"PER_MIN",
	"SLOWEST10PERCENT",
	"SLOWEST5PERCENT",
	"SUM",
	"SUM_DIMENSIONS",
}

// UnmarshalJSON performs custom unmarshalling of this enum Aggregation
func (t *Aggregation) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = Aggregation(name)
	return nil
}
