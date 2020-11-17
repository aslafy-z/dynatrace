package dashboards

// DashboardStub is a short representation of a dashboard
type DashboardStub struct {
	ID    string  `json:"id"`              // the ID of the dashboard
	Name  *string `json:"name,omitempty"`  // the name of the dashboard
	Owner *string `json:"owner,omitempty"` // the owner of the dashboard
}
