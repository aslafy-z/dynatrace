package storage

// TransactionTrafficQuota represetnts the maximum number of newly monitored entry point PurePaths captured per process/minute on environment level. Can be set to any value from 100 to 100000. If skipped when editing via PUT method then already set limit will remain
type TransactionTrafficQuota struct {
	MaxLimit *int32 `json:"maxLimit,omitempty"` // Maximum traffic [units per minute]
}
