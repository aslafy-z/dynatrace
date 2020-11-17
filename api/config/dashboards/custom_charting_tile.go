package dashboards

// CustomChartingTile Configuration of a custom chart tile
type CustomChartingTile struct {
	AbstractTile `json:",tileType=CUSTOM_CHARTING"`
	FilterConfig *CustomFilterConfig `json:"filterConfig,omitempty"`
}

// AsCustomChartingTile has not documentation
func (cct *CustomChartingTile) AsCustomChartingTile() *CustomChartingTile {
	return cct
}

// AsMarkdownTile has not documentation
func (cct *CustomChartingTile) AsMarkdownTile() *MarkdownTile {
	return nil
}

// AsUserSessionQueryTile has not documentation
func (cct *CustomChartingTile) AsUserSessionQueryTile() *UserSessionQueryTile {
	return nil
}

// AsFilterableEntityTile has not documentation
func (cct *CustomChartingTile) AsFilterableEntityTile() *FilterableEntityTile {
	return nil
}

// AsAssignedEntitiesWithMetricTile has not documentation
func (cct *CustomChartingTile) AsAssignedEntitiesWithMetricTile() *AssignedEntitiesWithMetricTile {
	return nil
}

// AsSyntheticSingleWebCheckTile has not documentation
func (cct *CustomChartingTile) AsSyntheticSingleWebCheckTile() *SyntheticSingleWebCheckTile {
	return nil
}

// AsAssignedEntitiesTile has not documentation
func (cct *CustomChartingTile) AsAssignedEntitiesTile() *AssignedEntitiesTile {
	return nil
}

// AsGenericTile has not documentation
func (cct *CustomChartingTile) AsGenericTile() *GenericTile {
	return nil
}
