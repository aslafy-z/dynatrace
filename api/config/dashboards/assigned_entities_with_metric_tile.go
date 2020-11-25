package dashboards

// AssignedEntitiesWithMetricTile Configuration of a tile with an assigned Dynatrace entity and an assigned metric. \n\nAn example is the Worldmap tile, showing the data from an assigned performance or behavior metric of an assigned application
type AssignedEntitiesWithMetricTile struct {
	AbstractTile     `json:",tileType=APPLICATION_WORLDMAP|RESOURCES|THIRD_PARTY_MOST_ACTIVE|UEM_CONVERSIONS_PER_GOAL,HOST|PROCESS_GROUPS_ONE"`
	AssignedEntities []string `json:"assignedEntities"` // The list of Dynatrace entities, assigned to the tile
	Metric           *string  `json:"metric,omitempty"` // The metric assigned to the tile
}

// AsCustomChartingTile has not documentation
func (aewmt *AssignedEntitiesWithMetricTile) AsCustomChartingTile() *CustomChartingTile {
	return nil
}

// AsMarkdownTile has not documentation
func (aewmt *AssignedEntitiesWithMetricTile) AsMarkdownTile() *MarkdownTile {
	return nil
}

// AsUserSessionQueryTile has not documentation
func (aewmt *AssignedEntitiesWithMetricTile) AsUserSessionQueryTile() *UserSessionQueryTile {
	return nil
}

// AsFilterableEntityTile has not documentation
func (aewmt *AssignedEntitiesWithMetricTile) AsFilterableEntityTile() *FilterableEntityTile {
	return nil
}

// AsAssignedEntitiesWithMetricTile has not documentation
func (aewmt *AssignedEntitiesWithMetricTile) AsAssignedEntitiesWithMetricTile() *AssignedEntitiesWithMetricTile {
	return aewmt
}

// AsSyntheticSingleWebCheckTile has not documentation
func (aewmt *AssignedEntitiesWithMetricTile) AsSyntheticSingleWebCheckTile() *SyntheticSingleWebCheckTile {
	return nil
}

// AsAssignedEntitiesTile has not documentation
func (aewmt *AssignedEntitiesWithMetricTile) AsAssignedEntitiesTile() *AssignedEntitiesTile {
	return nil
}

/*
// AsGenericTile has not documentation
func (aewmt *AssignedEntitiesWithMetricTile) AsGenericTile() *GenericTile {
	return nil
}
*/
