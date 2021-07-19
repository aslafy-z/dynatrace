package quota

// CustomMetrics represents custom metrics consumption and quota information on environment level. Not set (and not editable) if Custom metrics is not enabled. Not set (and not editable) if Davis data units is enabled. If skipped when editing via PUT method then already set quota will remain
type CustomMetrics struct {
	MaxLimit *int64 `json:"maxLimit,omitempty"` // Concurrent environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
	// READ ONLY
	CurrentUsage *float64 `json:"currentUsage,omitempty"` // Current environment usage
}
