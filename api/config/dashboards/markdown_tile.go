package dashboards

// MarkdownTile Configuration of the Markdown tile
type MarkdownTile struct {
	AbstractTile `json:",tileType=MARKDOWN"`
	Markdown     string `json:"markdown"` // The markdown-formatted content of the tile
}

// AsCustomChartingTile has not documentation
func (mdt *MarkdownTile) AsCustomChartingTile() *CustomChartingTile {
	return nil
}

// AsMarkdownTile has not documentation
func (mdt *MarkdownTile) AsMarkdownTile() *MarkdownTile {
	return mdt
}

// AsUserSessionQueryTile has not documentation
func (mdt *MarkdownTile) AsUserSessionQueryTile() *UserSessionQueryTile {
	return nil
}

// AsFilterableEntityTile has not documentation
func (mdt *MarkdownTile) AsFilterableEntityTile() *FilterableEntityTile {
	return nil
}

// AsAssignedEntitiesWithMetricTile has not documentation
func (mdt *MarkdownTile) AsAssignedEntitiesWithMetricTile() *AssignedEntitiesWithMetricTile {
	return nil
}

// AsSyntheticSingleWebCheckTile has not documentation
func (mdt *MarkdownTile) AsSyntheticSingleWebCheckTile() *SyntheticSingleWebCheckTile {
	return nil
}

// AsAssignedEntitiesTile has not documentation
func (mdt *MarkdownTile) AsAssignedEntitiesTile() *AssignedEntitiesTile {
	return nil
}

/*
// AsGenericTile has not documentation
func (mdt *MarkdownTile) AsGenericTile() *GenericTile {
	return nil
}
*/
