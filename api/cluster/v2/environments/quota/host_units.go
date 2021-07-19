package quota

// HostUnits represents host units consumption and quota information on environment level. If skipped when editing via PUT method then already set quota will remain
type HostUnits struct {
	MaxLimit *int64 `json:"maxLimit,omitempty"` // Concurrent environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
	// READ ONLY
	CurrentUsage *float64 `json:"currentUsage,omitempty"` // Current environment usage
}
