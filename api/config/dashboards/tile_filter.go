package dashboards

import (
	api "github.com/dtcookie/dynatrace/api/config"
)

// TileFilter is filter applied to a tile.
// It overrides dashboard's filter
type TileFilter struct {
	Timeframe      *string                        `json:"timeframe,omitempty"` // the default timeframe of the dashboard
	ManagementZone *api.EntityShortRepresentation `json:"managementZone,omitempty"`
}
