package dashboards

import "encoding/json"

// CustomFilterChartConfigType has no documentation
type CustomFilterChartConfigType string

// CustomFilterChartConfigTypes offers the known enum values
var CustomFilterChartConfigTypes = struct {
	Pie         CustomFilterChartConfigType
	SingleValue CustomFilterChartConfigType
	TimeSeries  CustomFilterChartConfigType
	TopList     CustomFilterChartConfigType
}{
	"PIE",
	"SINGLE_VALUE",
	"TIMESERIES",
	"TOP_LIST",
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *CustomFilterChartConfigType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = CustomFilterChartConfigType(name)
	return nil
}
