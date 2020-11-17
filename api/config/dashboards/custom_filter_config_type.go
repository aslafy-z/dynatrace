package dashboards

import "encoding/json"

// CustomFilterConfigType has no documentation
type CustomFilterConfigType string

// CustomFilterConfigTypes offers the known enum values
var CustomFilterConfigTypes = struct {
	Alb                       CustomFilterConfigType
	Application               CustomFilterConfigType
	ApplicationMethod         CustomFilterConfigType
	Appmon                    CustomFilterConfigType
	Asg                       CustomFilterConfigType
	AwsCredentials            CustomFilterConfigType
	AwsCustomService          CustomFilterConfigType
	AwsLambdaFunction         CustomFilterConfigType
	CloudApplication          CustomFilterConfigType
	CloudApplicationInstance  CustomFilterConfigType
	CloudApplicationNamespace CustomFilterConfigType
	ContainerGroupInstance    CustomFilterConfigType
	CustomApplication         CustomFilterConfigType
	CustomDevices             CustomFilterConfigType
	CustomServices            CustomFilterConfigType
	Database                  CustomFilterConfigType
	DatabaseKeyRequest        CustomFilterConfigType
	DcrumApplication          CustomFilterConfigType
	DcrumEntity               CustomFilterConfigType
	DynamoDb                  CustomFilterConfigType
	Ebs                       CustomFilterConfigType
	Ec2                       CustomFilterConfigType
	Elb                       CustomFilterConfigType
	Environment               CustomFilterConfigType
	Esxi                      CustomFilterConfigType
	ExternalSyntheticTest     CustomFilterConfigType
	GlobalBackgroundActivity  CustomFilterConfigType
	Host                      CustomFilterConfigType
	Iot                       CustomFilterConfigType
	KubernetesCluster         CustomFilterConfigType
	KubernetesNode            CustomFilterConfigType
	MdaService                CustomFilterConfigType
	Mixed                     CustomFilterConfigType
	MobileApplication         CustomFilterConfigType
	MonitoredEntity           CustomFilterConfigType
	Nlb                       CustomFilterConfigType
	PgBackgroundActivity      CustomFilterConfigType
	Problem                   CustomFilterConfigType
	ProcessGroupInstance      CustomFilterConfigType
	Rds                       CustomFilterConfigType
	RemotePlugin              CustomFilterConfigType
	Service                   CustomFilterConfigType
	ServiceKeyRequest         CustomFilterConfigType
	SyntheticBrowserMonitor   CustomFilterConfigType
	SyntheticHTTPcheck        CustomFilterConfigType
	SyntheticHTTPcheckStep    CustomFilterConfigType
	SyntheticLocation         CustomFilterConfigType
	SyntheticTest             CustomFilterConfigType
	SyntheticTestStep         CustomFilterConfigType
	UUEntity                  CustomFilterConfigType
	VirtualMachine            CustomFilterConfigType
	WebCheck                  CustomFilterConfigType
}{
	"ALB",
	"APPLICATION",
	"APPLICATION_METHOD",
	"APPMON",
	"ASG",
	"AWS_CREDENTIALS",
	"AWS_CUSTOM_SERVICE",
	"AWS_LAMBDA_FUNCTION",
	"CLOUD_APPLICATION",
	"CLOUD_APPLICATION_INSTANCE",
	"CLOUD_APPLICATION_NAMESPACE",
	"CONTAINER_GROUP_INSTANCE",
	"CUSTOM_APPLICATION",
	"CUSTOM_DEVICES",
	"CUSTOM_SERVICES",
	"DATABASE",
	"DATABASE_KEY_REQUEST",
	"DCRUM_APPLICATION",
	"DCRUM_ENTITY",
	"DYNAMO_DB",
	"EBS",
	"EC2",
	"ELB",
	"ENVIRONMENT",
	"ESXI",
	"EXTERNAL_SYNTHETIC_TEST",
	"GLOBAL_BACKGROUND_ACTIVITY",
	"HOST",
	"IOT",
	"KUBERNETES_CLUSTER",
	"KUBERNETES_NODE",
	"MDA_SERVICE",
	"MIXED",
	"MOBILE_APPLICATION",
	"MONITORED_ENTITY",
	"NLB",
	"PG_BACKGROUND_ACTIVITY",
	"PROBLEM",
	"PROCESS_GROUP_INSTANCE",
	"RDS",
	"REMOTE_PLUGIN",
	"SERVICE",
	"SERVICE_KEY_REQUEST",
	"SYNTHETIC_BROWSER_MONITOR",
	"SYNTHETIC_HTTPCHECK",
	"SYNTHETIC_HTTPCHECK_STEP",
	"SYNTHETIC_LOCATION",
	"SYNTHETIC_TEST",
	"SYNTHETIC_TEST_STEP",
	"UI_ENTITY",
	"VIRTUAL_MACHINE",
	"WEB_CHECK",
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *CustomFilterConfigType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = CustomFilterConfigType(name)
	return nil
}
