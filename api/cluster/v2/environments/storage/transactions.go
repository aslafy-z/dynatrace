package storage

// Transactions represents transaction storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain
type Transactions struct {
	MaxLimit *int64 `json:"maxLimit,omitempty"` // Maximum storage limit [bytes]
	// READ ONLY
	RetentionReductionPercentage *uint8  `json:"retentionReductionPercentage,omitempty"` // Percentage of truncation for new data
	RetentionReductionReason     *string `json:"retentionReductionReason,omitempty"`     // Reason of truncation
	CurrentlyUsed                *int64  `json:"currentlyUsed,omitempty"`                // Currently used storage [bytes]
}
