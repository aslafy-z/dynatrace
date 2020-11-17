package dashboards

import "encoding/json"

// CustomFilterChartSeriesConfigType has no documentation
type CustomFilterChartSeriesConfigType string

// CustomFilterChartSeriesConfigTypes offers the known enum values
var CustomFilterChartSeriesConfigTypes = struct {
	Area CustomFilterChartSeriesConfigType
	Bar  CustomFilterChartSeriesConfigType
	Line CustomFilterChartSeriesConfigType
}{
	"AREA",
	"BAR",
	"LINE",
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *CustomFilterChartSeriesConfigType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = CustomFilterChartSeriesConfigType(name)
	return nil
}
