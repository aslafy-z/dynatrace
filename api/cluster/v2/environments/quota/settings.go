package quota

// Settings represents environment level consumption and quotas information. Only returned if includeConsumptionInfo or includeUncachedConsumptionInfo param is true. If skipped when editing via PUT method then already set quotas will remain
type Settings struct {
	HostUnits     *HostUnits      `json:"hostUnits,omitempty"`         // Host units consumption and quota information on environment level. If skipped when editing via PUT method then already set quota will remain
	DEMUnits      *DEMUnits       `json:"demUnits,omitempty"`          // DEM units consumption and quota information on environment level. Not set (and not editable) if DEM units is not enabled. If skipped when editing via PUT method then already set quotas will remain
	UserSessions  *UserSessions   `json:"userSessions,omitempty"`      // User sessions consumption and quota information on environment level. If skipped when editing via PUT method then already set quotas will remain
	Synthetic     *Synthetic      `json:"syntheticMonitors,omitempty"` // Synthetic monitors consumption and quota information on environment level. Not set (and not editable) if neither Synthetic nor DEM units is enabled. If skipped when editing via PUT method then already set quotas will remain
	CustomMetrics *CustomMetrics  `json:"customMetrics,omitempty"`     // Custom metrics consumption and quota information on environment level. Not set (and not editable) if Custom metrics is not enabled. Not set (and not editable) if Davis data units is enabled. If skipped when editing via PUT method then already set quota will remain
	DDUs          *DavisDataUnits `json:"davisDataUnits,omitempty"`    // Davis Data Units consumption and quota information on environment level. Not set (and not editable) if Davis data units is not enabled. If skipped when editing via PUT method then already set quotas will remain
	LogMonitoring *LogMonitoring  `json:"logMonitoring,omitempty"`     // Log Monitoring consumption and quota information on environment level. Not set (and not editable) if Log monitoring is not enabled. Not set (and not editable) if Log monitoring is migrated to Davis data on license level. If skipped when editing via PUT method then already set quotas will remain
	// READ ONLY
	SessionProperties *SessionProperties `json:"sessionProperties,omitempty"` // User session properties consumption information on environment level

}
