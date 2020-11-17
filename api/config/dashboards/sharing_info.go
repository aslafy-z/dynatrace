package dashboards

// SharingInfo represents sharing configuration of a dashboard
type SharingInfo struct {
	LinkShared *bool `json:"linkShared,omitempty"` // If `true`, the dashboard is shared via link and authenticated users with the link can view
	Published  *bool `json:"published,omitempty"`  // If `true`, the dashboard is published to anyone on this environment
}
