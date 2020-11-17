package dashboards

// DashboardMetadata contains parameters of a dashboard
type DashboardMetadata struct {
	Name            string           `json:"name"`                     // the name of the dashboard
	Shared          *bool            `json:"shared,omitempty"`         // the dashboard is shared (`true`) or private (`false`)
	Owner           *string          `json:"owner,omitempty"`          // the owner of the dashboard
	SharingDetails  *SharingInfo     `json:"sharingDetails,omitempty"` // represents sharing configuration of a dashboard
	Filter          *DashboardFilter `json:"dashboardFilter,omitempty"`
	Tags            []string         `json:"tags,omitempty"`            // a set of tags assigned to the dashboard
	Preset          *bool            `json:"preset,omitempty"`          // the dashboard is a preset (`true`)
	ValidFilterKeys []string         `json:"validFilterKeys,omitempty"` // a set of all possible global dashboard filters that can be applied to dashboard
}
