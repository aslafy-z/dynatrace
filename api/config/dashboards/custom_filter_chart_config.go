package dashboards

// CustomFilterChartConfig Configuration of a custom chart
type CustomFilterChartConfig struct {
	LegendShown         *bool                                       `json:"legendShown,omitempty"` // Defines if a legend should be shown
	Type                CustomFilterChartConfigType                 `json:"type"`                  // The type of the chart
	Series              []CustomFilterChartSeriesConfig             `json:"series"`                // A list of charted metrics
	ResultMetadata      map[string]CustomChartingItemMetadataConfig `json:"resultMetadata"`        // Additional information about charted metric
	AxisLimits          map[string]float64                          `json:"axisLimits,omitempty"`  // The optional custom y-axis limits
	LeftAxisCustomUnit  *LeftAxisCustomUnit                         `json:"leftAxisCustomUnit,omitempty"`
	RightAxisCustomUnit *RightAxisCustomUnit                        `json:"rightAxisCustomUnit,omitempty"`
}
