package dashboards

// DashboardList is a list of short representations of dashboards
type DashboardList struct {
	Dashboards []DashboardStub `json:"dashboards,omitempty"` // the short representations of the dashboards
}
