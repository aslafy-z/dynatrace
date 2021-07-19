package retention

// RealUserMonitoring represents Real user monitoring retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain
type RealUserMonitoring struct {
	MaxLimitInDays *int64 `json:"maxLimitInDays,omitempty"` // Maximum retention limit [days]
	// READ ONLY
	CurrentlyUsedInMillis *int64 `json:"currentlyUsedInMillis,omitempty"` // Current data age [milliseconds]
	CurrentlyUsedInDays   *int64 `json:"currentlyUsedInDays,omitempty"`   // Current data age [days]
}
