package dashboards

// CustomFilterConfig Configuration of the custom filter of a tile
type CustomFilterConfig struct {
	Type                 CustomFilterConfigType         `json:"type"`                 // The type of the filter
	CustomName           string                         `json:"customName"`           // The name of the tile, set by user
	DefaultName          string                         `json:"defaultName"`          // The default name of the tile
	ChartConfig          *CustomFilterChartConfig       `json:"chartConfig"`          // Config Configuration of a custom chart
	FiltersPerEntityType map[string]map[string][]string `json:"filtersPerEntityType"` // A list of filters, applied to specific entity types
}
