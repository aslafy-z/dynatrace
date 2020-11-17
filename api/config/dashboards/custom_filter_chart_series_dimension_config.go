package dashboards

// CustomFilterChartSeriesDimensionConfig Configuration of the charted metric splitting
type CustomFilterChartSeriesDimensionConfig struct {
	ID              string   `json:"id"`             // The ID of the dimension by which the metric is split
	Name            *string  `json:"name,omitempty"` // The name of the dimension by which the metric is split
	Values          []string `json:"values"`         // The splitting value
	EntityDimension *bool    `json:"entityDimension,omitempty"`
}
