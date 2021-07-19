package sharing

import "github.com/dtcookie/hcl"

// DashboardSharing represents sharing configuration of the dashboard
type DashboardSharing struct {
	DashboardID  string           `json:"id"`           // The Dynatrace entity ID of the dashboard
	Permissions  SharePermissions `json:"permissions"`  // Access permissions of the dashboard
	PublicAccess *AnonymousAccess `json:"publicAccess"` // Configuration of the [anonymous access](https://dt-url.net/ov03sf1) to the dashboard
	Preset       bool             `json:"preset"`       // If `true` the dashboard will be marked as preset
	Enabled      bool             `json:"enabled"`      // The dashboard is shared (`true`) or private (`false`)
}

func (me *DashboardSharing) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dashboard_id": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The Dynatrace entity ID of the dashboard",
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "The dashboard is shared (`true`) or private (`false`)",
		},
		"preset": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "If `true` the dashboard will be marked as preset",
		},
		"permissions": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        hcl.Resource{Schema: new(SharePermissions).Schema()},
			Description: "Access permissions of the dashboard",
		},
		"public": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        hcl.Resource{Schema: new(AnonymousAccess).Schema()},
			Description: "Configuration of the [anonymous access](https://dt-url.net/ov03sf1) to the dashboard",
		},
	}
}
