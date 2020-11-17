package managementzones

// ComparisonBasicType Defines the actual set of fields depending on the value. See one of the following objects:
// * `STRING` -> StringComparison
// * `INDEXED_NAME` -> IndexedNameComparison
// * `INDEXED_STRING` -> IndexedStringComparison
// * `INTEGER` -> IntegerComparison
// * `SERVICE_TYPE` -> ServiceTypeComparison
// * `PAAS_TYPE` -> PaasTypeComparison
// * `CLOUD_TYPE` -> CloudTypeComparison
// * `AZURE_SKU` -> AzureSkuComparision
// * `AZURE_COMPUTE_MODE` -> AzureComputeModeComparison
// * `ENTITY_ID` -> EntityIdComparison
// * `SIMPLE_TECH` -> SimpleTechComparison
// * `SIMPLE_HOST_TECH` -> SimpleHostTechComparison
// * `SERVICE_TOPOLOGY` -> ServiceTopologyComparison
// * `DATABASE_TOPOLOGY` -> DatabaseTopologyComparison
// * `OS_TYPE` -> OsTypeComparison
// * `HYPERVISOR_TYPE` -> HypervisorTypeComparision
// * `IP_ADDRESS` -> IpAddressComparison
// * `OS_ARCHITECTURE` -> OsArchitectureComparison
// * `BITNESS` -> BitnessComparision
// * `APPLICATION_TYPE` -> ApplicationTypeComparison
// * `MOBILE_PLATFORM` -> MobilePlatformComparison
// * `CUSTOM_APPLICATION_TYPE` -> CustomApplicationTypeComparison
// * `DCRUM_DECODER_TYPE` -> DcrumDecoderComparison
// * `SYNTHETIC_ENGINE_TYPE` -> SyntheticEngineTypeComparison
// * `TAG` -> TagComparison
// * `INDEXED_TAG` -> IndexedTagComparison
type ComparisonBasicType string

// ComparisonBasicTypes offers the known enum values
var ComparisonBasicTypes = struct {
	ApplicationType       ComparisonBasicType
	AzureComputeMode      ComparisonBasicType
	AzureSku              ComparisonBasicType
	Bitness               ComparisonBasicType
	CloudType             ComparisonBasicType
	CustomApplicationType ComparisonBasicType
	DatabaseTopology      ComparisonBasicType
	DCRumDecoderType      ComparisonBasicType
	EntityID              ComparisonBasicType
	HypervisorType        ComparisonBasicType
	IndexedName           ComparisonBasicType
	IndexedString         ComparisonBasicType
	IndexedTag            ComparisonBasicType
	Integer               ComparisonBasicType
	IPAddress             ComparisonBasicType
	MobilePlatform        ComparisonBasicType
	OSArchitecture        ComparisonBasicType
	OSType                ComparisonBasicType
	PaasType              ComparisonBasicType
	ServiceTopology       ComparisonBasicType
	ServiceType           ComparisonBasicType
	SimpleHostTech        ComparisonBasicType
	SimpleTech            ComparisonBasicType
	String                ComparisonBasicType
	SyntheticEngineType   ComparisonBasicType
	Tag                   ComparisonBasicType
}{
	"APPLICATION_TYPE",
	"AZURE_COMPUTE_MODE",
	"AZURE_SKU",
	"BITNESS",
	"CLOUD_TYPE",
	"CUSTOM_APPLICATION_TYPE",
	"DATABASE_TOPOLOGY",
	"DCRUM_DECODER_TYPE",
	"ENTITY_ID",
	"HYPERVISOR_TYPE",
	"INDEXED_NAME",
	"INDEXED_STRING",
	"INDEXED_TAG",
	"INTEGER",
	"IP_ADDRESS",
	"MOBILE_PLATFORM",
	"OS_ARCHITECTURE",
	"OS_TYPE",
	"PAAS_TYPE",
	"SERVICE_TOPOLOGY",
	"SERVICE_TYPE",
	"SIMPLE_HOST_TECH",
	"SIMPLE_TECH",
	"STRING",
	"SYNTHETIC_ENGINE_TYPE",
	"TAG",
}
