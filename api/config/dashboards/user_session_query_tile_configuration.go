package dashboards

// UserSessionQueryTileConfiguration Configuration of a User session query visualization tile
type UserSessionQueryTileConfiguration struct {
	HasAxisBucketing *bool `json:"hasAxisBucketing,omitempty"` // The axis bucketing when enabled groups similar series in the same virtual axis
}
