package managementzones

// ManagementZoneRuleType The type of Dynatrace entities the management zone can be applied to.
type ManagementZoneRuleType string

// ManagementZoneRuleTypes offers the known enum values
var ManagementZoneRuleTypes = struct {
	AppMonServer                 ManagementZoneRuleType
	AppMonSystemProfile          ManagementZoneRuleType
	AWSAccount                   ManagementZoneRuleType
	AWSApplicationLoadBalancer   ManagementZoneRuleType
	AWSAutoScalingGroup          ManagementZoneRuleType
	AWSClassicLoadBalancer       ManagementZoneRuleType
	AWSNetworkLoadBalancer       ManagementZoneRuleType
	AWSRelationalDatabaseService ManagementZoneRuleType
	Azure                        ManagementZoneRuleType
	BrowserMonitor               ManagementZoneRuleType
	CloudApplication             ManagementZoneRuleType
	CloudApplicationNamespace    ManagementZoneRuleType
	CloudFoundryFoundation       ManagementZoneRuleType
	CustomApplication            ManagementZoneRuleType
	CustomDevice                 ManagementZoneRuleType
	CustomDeviceGroup            ManagementZoneRuleType
	DataCenterService            ManagementZoneRuleType
	EnterpriseApplication        ManagementZoneRuleType
	ESXIHost                     ManagementZoneRuleType
	ExternalMonitor              ManagementZoneRuleType
	Host                         ManagementZoneRuleType
	HostGroup                    ManagementZoneRuleType
	HTTPMonitor                  ManagementZoneRuleType
	KubernetesCluster            ManagementZoneRuleType
	MobileApplication            ManagementZoneRuleType
	OpenStackAccount             ManagementZoneRuleType
	ProcessGroup                 ManagementZoneRuleType
	Service                      ManagementZoneRuleType
	WebApplication               ManagementZoneRuleType
}{
	"APPMON_SERVER",
	"APPMON_SYSTEM_PROFILE",
	"AWS_ACCOUNT",
	"AWS_APPLICATION_LOAD_BALANCER",
	"AWS_AUTO_SCALING_GROUP",
	"AWS_CLASSIC_LOAD_BALANCER",
	"AWS_NETWORK_LOAD_BALANCER",
	"AWS_RELATIONAL_DATABASE_SERVICE",
	"AZURE",
	"BROWSER_MONITOR",
	"CLOUD_APPLICATION",
	"CLOUD_APPLICATION_NAMESPACE",
	"CLOUD_FOUNDRY_FOUNDATION",
	"CUSTOM_APPLICATION",
	"CUSTOM_DEVICE",
	"CUSTOM_DEVICE_GROUP",
	"DATA_CENTER_SERVICE",
	"ENTERPRISE_APPLICATION",
	"ESXI_HOST",
	"EXTERNAL_MONITOR",
	"HOST",
	"HOST_GROUP",
	"HTTP_MONITOR",
	"KUBERNETES_CLUSTER",
	"MOBILE_APPLICATION",
	"OPENSTACK_ACCOUNT",
	"PROCESS_GROUP",
	"SERVICE",
	"WEB_APPLICATION",
}
