package managementzones

// PaasTypeComparisonValue The value to compare to.
type PaasTypeComparisonValue string

// PaasTypeComparisonValues offers the known enum values
var PaasTypeComparisonValues = struct {
	AWSECSEC2       PaasTypeComparisonValue
	AWSECSFargate   PaasTypeComparisonValue
	AWSLambda       PaasTypeComparisonValue
	AzureFunctions  PaasTypeComparisonValue
	AzureWebsites   PaasTypeComparisonValue
	CloudFoundry    PaasTypeComparisonValue
	GoogleAppEngine PaasTypeComparisonValue
	Heroku          PaasTypeComparisonValue
	Kubernetes      PaasTypeComparisonValue
	Openshift       PaasTypeComparisonValue
}{
	"AWS_ECS_EC2",
	"AWS_ECS_FARGATE",
	"AWS_LAMBDA",
	"AZURE_FUNCTIONS",
	"AZURE_WEBSITES",
	"CLOUD_FOUNDRY",
	"GOOGLE_APP_ENGINE",
	"HEROKU",
	"KUBERNETES",
	"OPENSHIFT",
}
