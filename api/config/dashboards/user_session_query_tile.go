package dashboards

// UserSessionQueryTile Configuration of a User session query tile
type UserSessionQueryTile struct {
	AbstractTile        `json:",tileType=DTAQL"`
	CustomName          string                             `json:"customName"`                    // The name of the tile, set by user
	Query               string                             `json:"query"`                         // A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile
	Type                UserSessionQueryTileType           `json:"type"`                          // The visualization of the tile
	TimeFrameShift      *string                            `json:"timeFrameShift,omitempty"`      // The comparison timeframe of the query. If specified, you additionally get the results of the same query with the specified time shift
	VisualizationConfig *UserSessionQueryTileConfiguration `json:"visualizationConfig,omitempty"` // Configuration of a User session query visualization tile
	Limit               *int32                             `json:"limit,omitempty"`               // The limit of the results, if not set will use the default value of the system
}

// AsCustomChartingTile has not documentation
func (usqt *UserSessionQueryTile) AsCustomChartingTile() *CustomChartingTile {
	return nil
}

// AsMarkdownTile has not documentation
func (usqt *UserSessionQueryTile) AsMarkdownTile() *MarkdownTile {
	return nil
}

// AsUserSessionQueryTile has not documentation
func (usqt *UserSessionQueryTile) AsUserSessionQueryTile() *UserSessionQueryTile {
	return usqt
}

// AsFilterableEntityTile has not documentation
func (usqt *UserSessionQueryTile) AsFilterableEntityTile() *FilterableEntityTile {
	return nil
}

// AsAssignedEntitiesWithMetricTile has not documentation
func (usqt *UserSessionQueryTile) AsAssignedEntitiesWithMetricTile() *AssignedEntitiesWithMetricTile {
	return nil
}

// AsSyntheticSingleWebCheckTile has not documentation
func (usqt *UserSessionQueryTile) AsSyntheticSingleWebCheckTile() *SyntheticSingleWebCheckTile {
	return nil
}

// AsAssignedEntitiesTile has not documentation
func (usqt *UserSessionQueryTile) AsAssignedEntitiesTile() *AssignedEntitiesTile {
	return nil
}

/*
// AsGenericTile has not documentation
func (usqt *UserSessionQueryTile) AsGenericTile() *GenericTile {
	return nil
}
*/
