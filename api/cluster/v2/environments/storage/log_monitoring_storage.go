package storage

// LogMonitoringStorage represents log monitoring storage usage and limit information on environment level. Not editable when Log monitoring is not allowed by license or not configured on cluster level. If skipped when editing via PUT method then already set limit will remain
type LogMonitoringStorage struct {
	MaxLimit *int64 `json:"maxLimit,omitempty"` // Maximum storage limit [bytes]
	// READ ONLY
	CurrentlyUsed *int64 `json:"currentlyUsed,omitempty"` // Currently used storage [bytes]
}
