package quota

// SessionProperties represents user session properties consumption information on environment level
type SessionProperties struct {
	// READ ONLY
	ConsumedThisMonth *float64 `json:"consumedThisMonth,omitempty"` // Monthly environment consumption. Resets each calendar month
	ConsumedThisYear  *float64 `json:"consumedThisYear,omitempty"`  // Yearly environment consumption. Resets each year on license creation date anniversary
}
