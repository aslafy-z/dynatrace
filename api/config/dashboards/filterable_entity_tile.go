package dashboards

// FilterableEntityTile Configuration of a tile with the built-in custom filter.
// An example is the Service health tile, which may use a custom timeframe
type FilterableEntityTile struct {
	AbstractTile `json:",tileType=HOSTS|APPLICATIONS|SERVICES|DATABASES_OVERVIEW|SYNTHETIC_TESTS"`
	FilterConfig *CustomFilterConfig `json:"filterConfig,omitempty"` // Configuration of the custom filter of a tile
	ChartVisible *bool               `json:"chartVisible,omitempty"`
}

// AsCustomChartingTile has not documentation
func (fet *FilterableEntityTile) AsCustomChartingTile() *CustomChartingTile {
	return nil
}

// AsMarkdownTile has not documentation
func (fet *FilterableEntityTile) AsMarkdownTile() *MarkdownTile {
	return nil
}

// AsUserSessionQueryTile has not documentation
func (fet *FilterableEntityTile) AsUserSessionQueryTile() *UserSessionQueryTile {
	return nil
}

// AsFilterableEntityTile has not documentation
func (fet *FilterableEntityTile) AsFilterableEntityTile() *FilterableEntityTile {
	return fet
}

// AsAssignedEntitiesWithMetricTile has not documentation
func (fet *FilterableEntityTile) AsAssignedEntitiesWithMetricTile() *AssignedEntitiesWithMetricTile {
	return nil
}

// AsSyntheticSingleWebCheckTile has not documentation
func (fet *FilterableEntityTile) AsSyntheticSingleWebCheckTile() *SyntheticSingleWebCheckTile {
	return nil
}

// AsAssignedEntitiesTile has not documentation
func (fet *FilterableEntityTile) AsAssignedEntitiesTile() *AssignedEntitiesTile {
	return nil
}

/*
// AsGenericTile has not documentation
func (fet *FilterableEntityTile) AsGenericTile() *GenericTile {
	return nil
}
*/
