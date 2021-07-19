package storage

import (
	"github.com/dtcookie/dynatrace/api/cluster/v2/environments/storage/retention"
	"github.com/dtcookie/hcl"
)

// Settings represents environment level storage usage and limit information. Not returned if includeStorageInfo param is not true. If skipped when editing via PUT method then already set limits will remain
type Settings struct {
	Transactions                 *Transactions                  `json:"transactionStorage,omitempty"`           // Transaction storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain
	SessionReplayStorage         *SessionReplayStorage          `json:"sessionReplayStorage,omitempty"`         // Session replay storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain
	SymbolFilesFromMobileApps    *SymbolFilesFromMobileApps     `json:"symbolFilesFromMobileApps,omitempty"`    // Symbol files from mobile apps storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain
	LogMonitoringStorage         *LogMonitoringStorage          `json:"logMonitoringStorage,omitempty"`         // Log monitoring storage usage and limit information on environment level. Not editable when Log monitoring is not allowed by license or not configured on cluster level. If skipped when editing via PUT method then already set limit will remain
	ServiceRequestLevelRetention *retention.ServiceRequestLevel `json:"serviceRequestLevelRetention,omitempty"` // Service request level retention settings on environment level. Service code level retention time can't be greater than service request level retention time and both can't exceed one year.If skipped when editing via PUT method then already set limit will remain
	ServiceCodeLevelRetention    *retention.ServiceCodeLevel    `json:"serviceCodeLevelRetention,omitempty"`    // Service code level retention settings on environment level. Service code level retention time can't be greater than service request level retention time and both can't exceed one year.If skipped when editing via PUT method then already set limit will remain
	RealUserMonitoringRetention  *retention.RealUserMonitoring  `json:"realUserMonitoringRetention,omitempty"`  // Real user monitoring retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain
	SyntheticMonitoringRetention *retention.SyntheticMonitoring `json:"syntheticMonitoringRetention,omitempty"` // Synthetic monitoring retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain
	SessionReplayRetention       *retention.SessionReplay       `json:"sessionReplayRetention,omitempty"`       // Session replay retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain
	LogMonitoringRetention       *retention.LogMonitoring       `json:"logMonitoringRetention,omitempty"`       // Log monitoring retention settings on environment level. Not editable when Log monitoring is not allowed by license or not configured on cluster level. Can be set to any value from 5 to 90 days. If skipped when editing via PUT method then already set limit will remain
	UserActionsPerMinute         *UserActionsPerMinute          `json:"userActionsPerMinute,omitempty"`         // Maximum number of user actions generated per minute on environment level. Can be set to any value from 1 to 2147483646 or left unlimited. If skipped when editing via PUT method then already set limit will remain
	TransactionTrafficQuota      *TransactionTrafficQuota       `json:"transactionTrafficQuota,omitempty"`      // Maximum number of newly monitored entry point PurePaths captured per process/minute on environment level. Can be set to any value from 100 to 100000. If skipped when editing via PUT method then already set limit will remain
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"transactions": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Transaction storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain",
		},
	}
}
