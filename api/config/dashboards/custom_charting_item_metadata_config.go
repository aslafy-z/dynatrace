package dashboards

// CustomChartingItemMetadataConfig Additional metadata for charted metric
type CustomChartingItemMetadataConfig struct {
	LastModified *int64 `json:"lastModified,omitempty"` // The timestamp of the last metadata modification, in UTC milliseconds
	CustomColor  string `json:"customColor"`            // The color of the metric in the chart, hex format
}
