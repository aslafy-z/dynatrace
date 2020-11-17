package dashboards

// CustomFilterChartSeriesConfig is the configuration of a charted metric
type CustomFilterChartSeriesConfig struct {
	Metric          string                                   `json:"metric"`               // The name of the charted metric
	Aggregation     Aggregation                              `json:"aggregation"`          // The charted aggregation of the metric
	Percentile      *int64                                   `json:"percentile,omitempty"` // The charted percentile. Only applicable if the **aggregation** is set to `PERCENTILE`.
	Type            CustomFilterChartSeriesConfigType        `json:"type"`                 // The visualization of the timeseries chart
	EntityType      string                                   `json:"entityType"`           // The type of the Dynatrace entity that delivered the charted metric
	Dimensions      []CustomFilterChartSeriesDimensionConfig `json:"dimensions"`           // Configuration of the charted metric splitting
	SortAscending   bool                                     `json:"sortAscending"`        // Sort ascending (`true`) or descending (`false`)
	SortColumn      bool                                     `json:"sortColumn"`
	AggregationRate *AggregationRate                         `json:"aggregationRate,omitempty"`
}
