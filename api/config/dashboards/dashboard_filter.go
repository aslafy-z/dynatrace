package dashboards

import api "github.com/dtcookie/dynatrace/api/config"

// DashboardFilter represents filters, applied to a dashboard
type DashboardFilter struct {
	Timeframe      *string                        `json:"timeframe,omitempty"` // the default timeframe of the dashboard
	ManagementZone *api.EntityShortRepresentation `json:"managementZone,omitempty"`
}
