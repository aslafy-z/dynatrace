package quota

// UserSessions represents user sessions consumption and quota information on environment level. If skipped when editing via PUT method then already set quotas will remain
type UserSessions struct {
	TotalAnnualLimit  *int64 `json:"totalAnnualLimit,omitempty"`  // Annual total User sessions environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
	TotalMonthlyLimit *int64 `json:"totalMonthlyLimit,omitempty"` // Monthly total User sessions environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited

	// READ ONLY
	TotalConsumedThisYear                                *float64 `json:"totalConsumedThisYear,omitempty"`                                // Yearly total User sessions environment consumption. Resets each year on license creation date anniversary
	TotalConsumedThisMonth                               *float64 `json:"totalConsumedThisMonth,omitempty"`                               // Monthly total User sessions environment consumption. Resets each calendar month
	ConsumedMobileSessionsThisMonth                      *float64 `json:"consumedMobileSessionsThisMonth,omitempty"`                      // Monthly Mobile user sessions environment consumption. Resets each calendar month
	ConsumedMobileSessionsThisYear                       *float64 `json:"consumedMobileSessionsThisYear,omitempty"`                       // Yearly Mobile user sessions environment consumption. Resets each year on license creation date anniversary
	ConsumedUserSessionsWithMobileSessionReplayThisYear  *float64 `json:"consumedUserSessionsWithMobileSessionReplayThisYear,omitempty"`  // Yearly Mobile user sessions with replay environment consumption. Resets each year on license creation date anniversary
	ConsumedUserSessionsWithWebSessionReplayThisMonth    *float64 `json:"consumedUserSessionsWithWebSessionReplayThisMonth,omitempty"`    // Monthly Web user sessions with replay environment consumption. Resets each calendar month
	ConsumedUserSessionsWithWebSessionReplayThisYear     *float64 `json:"consumedUserSessionsWithWebSessionReplayThisYear,omitempty"`     // Yearly Web user sessions with replay environment consumption. Resets each year on license creation date anniversary
	ConsumedUserSessionsWithMobileSessionReplayThisMonth *float64 `json:"consumedUserSessionsWithMobileSessionReplayThisMonth,omitempty"` // Monthly Mobile user sessions with replay environment consumption. Resets each calendar month
}
