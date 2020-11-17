package dashboards

import (
	api "github.com/dtcookie/dynatrace/api/config"
)

// Dashboard the configuration of a dashboard
type Dashboard struct {
	ConfigurationMetadata *api.ConfigurationMetadata `json:"metadata,omitempty"`

	ID       *string            `json:"id,omitempty"`      // the ID of the dashboard
	Metadata *DashboardMetadata `json:"dashboardMetadata"` // contains parameters of a dashboard
	Tiles    TileCollection     `json:"tiles"`             // the tiles the dashboard consists of
}
