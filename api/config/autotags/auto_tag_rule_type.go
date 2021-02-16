package autotags

// AutoTagRuleType Type of entities to which the rule applies.
type AutoTagRuleType string

// AutoTagRuleTypes offers the known enum values
var AutoTagRuleTypes = struct {
	Application                  AutoTagRuleType
	AWSApplicationLoadBalancer   AutoTagRuleType
	AWSClassicLoadBalancer       AutoTagRuleType
	AWSNetworkLoadBalancer       AutoTagRuleType
	AWSRelationalDatabaseService AutoTagRuleType
	Azure                        AutoTagRuleType
	CustomApplication            AutoTagRuleType
	CustomDevice                 AutoTagRuleType
	DCRumApplication             AutoTagRuleType
	ESXIHost                     AutoTagRuleType
	ExternalSyntheticTest        AutoTagRuleType
	Host                         AutoTagRuleType
	HTTPCheck                    AutoTagRuleType
	MobileApplication            AutoTagRuleType
	ProcessGroup                 AutoTagRuleType
	Service                      AutoTagRuleType
	SyntheticTest                AutoTagRuleType
}{
	"APPLICATION",
	"AWS_APPLICATION_LOAD_BALANCER",
	"AWS_CLASSIC_LOAD_BALANCER",
	"AWS_NETWORK_LOAD_BALANCER",
	"AWS_RELATIONAL_DATABASE_SERVICE",
	"AZURE",
	"CUSTOM_APPLICATION",
	"CUSTOM_DEVICE",
	"DCRUM_APPLICATION",
	"ESXI_HOST",
	"EXTERNAL_SYNTHETIC_TEST",
	"HOST",
	"HTTP_CHECK",
	"MOBILE_APPLICATION",
	"PROCESS_GROUP",
	"SERVICE",
	"SYNTHETIC_TEST",
}
