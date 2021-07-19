package storage

// SymbolFilesFromMobileApps represents symbol files from mobile apps storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain
type SymbolFilesFromMobileApps struct {
	MaxLimit *int64 `json:"maxLimit,omitempty"` // Maximum storage limit [bytes]
	// READ ONLY
	CurrentlyUsed *int64 `json:"currentlyUsed,omitempty"` // Currently used storage [bytes]
}
