package dashboards

import "encoding/json"

// AggregationRate has no documentation
type AggregationRate string

// AggregationRates offers the known enum values
var AggregationRates = struct {
	Hour   AggregationRate
	Minute AggregationRate
	Second AggregationRate
	Total  AggregationRate
}{
	"HOUR",
	"MINUTE",
	"SECOND",
	"TOTAL",
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *AggregationRate) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = AggregationRate(name)
	return nil
}
