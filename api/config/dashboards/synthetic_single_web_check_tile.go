package dashboards

// SyntheticSingleWebCheckTile Configuration of the Browser monitor tile
type SyntheticSingleWebCheckTile struct {
	AbstractTile              `json:",tileType=SYNTHETIC_SINGLE_WEBCHECK"`
	AssignedEntities          []string `json:"assignedEntities"`                    // The list of Dynatrace entities, assigned to the tile
	ExcludeMaintenanceWindows *bool    `json:"excludeMaintenanceWindows,omitempty"` // Include (`false') or exclude (`true`) maintenance windows from availability calculations
}

// AsCustomChartingTile has not documentation
func (sswct *SyntheticSingleWebCheckTile) AsCustomChartingTile() *CustomChartingTile {
	return nil
}

// AsMarkdownTile has not documentation
func (sswct *SyntheticSingleWebCheckTile) AsMarkdownTile() *MarkdownTile {
	return nil
}

// AsUserSessionQueryTile has not documentation
func (sswct *SyntheticSingleWebCheckTile) AsUserSessionQueryTile() *UserSessionQueryTile {
	return nil
}

// AsFilterableEntityTile has not documentation
func (sswct *SyntheticSingleWebCheckTile) AsFilterableEntityTile() *FilterableEntityTile {
	return nil
}

// AsAssignedEntitiesWithMetricTile has not documentation
func (sswct *SyntheticSingleWebCheckTile) AsAssignedEntitiesWithMetricTile() *AssignedEntitiesWithMetricTile {
	return nil
}

// AsSyntheticSingleWebCheckTile has not documentation
func (sswct *SyntheticSingleWebCheckTile) AsSyntheticSingleWebCheckTile() *SyntheticSingleWebCheckTile {
	return sswct
}

// AsAssignedEntitiesTile has not documentation
func (sswct *SyntheticSingleWebCheckTile) AsAssignedEntitiesTile() *AssignedEntitiesTile {
	return nil
}

/*
// AsGenericTile has not documentation
func (sswct *SyntheticSingleWebCheckTile) AsGenericTile() *GenericTile {
	return nil
}
*/
