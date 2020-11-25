package dashboards

// AssignedEntitiesTile Configuration of a tile with an assigned Dynatrace entity. An example is the Bounce rate tile, showing the data from an assigned application.
type AssignedEntitiesTile struct {
	AbstractTile     `json:",tileType=VIRTUALIZATION|APPLICATION|AWS|SERVICE_VERSATILE|SESSION_METRICS|USERS|UEM_KEY_USER_ACTIONS|BOUNCE_RATE|UEM_CONVERSIONS_OVERALL|SYNTHETIC_HTTP_MONITOR|DATABASE|CUSTOM_APPLICATION|APPLICATION_METHOD|LOG_ANALYTICS|OPENSTACK|OPENSTACK_PROJECT|OPENSTACK_AV_ZONE|DEVICE_APPLICATION_METHOD|DEM_KEY_USER_ACTION"`
	AssignedEntities []string `json:"assignedEntities"` // The list of Dynatrace entities, assigned to the tile
}

// AsCustomChartingTile has not documentation
func (aet *AssignedEntitiesTile) AsCustomChartingTile() *CustomChartingTile {
	return nil
}

// AsMarkdownTile has not documentation
func (aet *AssignedEntitiesTile) AsMarkdownTile() *MarkdownTile {
	return nil
}

// AsUserSessionQueryTile has not documentation
func (aet *AssignedEntitiesTile) AsUserSessionQueryTile() *UserSessionQueryTile {
	return nil
}

// AsFilterableEntityTile has not documentation
func (aet *AssignedEntitiesTile) AsFilterableEntityTile() *FilterableEntityTile {
	return nil
}

// AsAssignedEntitiesWithMetricTile has not documentation
func (aet *AssignedEntitiesTile) AsAssignedEntitiesWithMetricTile() *AssignedEntitiesWithMetricTile {
	return nil
}

// AsSyntheticSingleWebCheckTile has not documentation
func (aet *AssignedEntitiesTile) AsSyntheticSingleWebCheckTile() *SyntheticSingleWebCheckTile {
	return nil
}

// AsAssignedEntitiesTile has not documentation
func (aet *AssignedEntitiesTile) AsAssignedEntitiesTile() *AssignedEntitiesTile {
	return aet
}

/*
// AsGenericTile has not documentation
func (aet *AssignedEntitiesTile) AsGenericTile() *GenericTile {
	return nil
}
*/
