package managementzones

import "github.com/dtcookie/gojson"

// ComparisonBasic Defines how the matching is actually performed: what and how are we comparing.
// The actual set of fields and possible values of the **operator** field depend on the **type** of the comparison. \n\nFind the list of actual models in the description of the **type** field and check the description of the model you need.
type ComparisonBasic interface {
	Initialize(*BaseComparisonBasic) // Initialize duplicates the properties of the other AbstractConditionKey into this instance. It also serves as a unique method for structs implementing the interface this abstract class is derived from
	Implementors() []ComparisonBasic // Implementors provides automatic JSON unmarshalling with a list of struct implementing (directly or indirectly) ConditionKey
}

// BaseComparisonBasic Defines how the matching is actually performed: what and how are we comparing.
// The actual set of fields and possible values of the **operator** field depend on the **type** of the comparison. \n\nFind the list of actual models in the description of the **type** field and check the description of the model you need.
type BaseComparisonBasic struct {
	Type   ComparisonBasicType `json:"type"`   // Defines the actual set of fields depending on the value. See one of the following objects:  * `STRING` -> StringComparison  * `INDEXED_NAME` -> IndexedNameComparison  * `INDEXED_STRING` -> IndexedStringComparison  * `INTEGER` -> IntegerComparison  * `SERVICE_TYPE` -> ServiceTypeComparison  * `PAAS_TYPE` -> PaasTypeComparison  * `CLOUD_TYPE` -> CloudTypeComparison  * `AZURE_SKU` -> AzureSkuComparision  * `AZURE_COMPUTE_MODE` -> AzureComputeModeComparison  * `ENTITY_ID` -> EntityIdComparison  * `SIMPLE_TECH` -> SimpleTechComparison  * `SIMPLE_HOST_TECH` -> SimpleHostTechComparison  * `SERVICE_TOPOLOGY` -> ServiceTopologyComparison  * `DATABASE_TOPOLOGY` -> DatabaseTopologyComparison  * `OS_TYPE` -> OsTypeComparison  * `HYPERVISOR_TYPE` -> HypervisorTypeComparision  * `IP_ADDRESS` -> IpAddressComparison  * `OS_ARCHITECTURE` -> OsArchitectureComparison  * `BITNESS` -> BitnessComparision  * `APPLICATION_TYPE` -> ApplicationTypeComparison  * `MOBILE_PLATFORM` -> MobilePlatformComparison  * `CUSTOM_APPLICATION_TYPE` -> CustomApplicationTypeComparison  * `DCRUM_DECODER_TYPE` -> DcrumDecoderComparison  * `SYNTHETIC_ENGINE_TYPE` -> SyntheticEngineTypeComparison  * `TAG` -> TagComparison  * `INDEXED_TAG` -> IndexedTagComparison
	Negate bool                `json:"negate"` // Reverses the comparison **operator**. For example it turns the **begins with** into **does not begin with**.
}

// Initialize duplicates the properties of the other BaseComparisonBasic into this instance
// It also serves as a unique method for structs implementing the interface this abstract class is derived from
func (bcb *BaseComparisonBasic) Initialize(other *BaseComparisonBasic) {
	bcb.Type = other.Type
	bcb.Negate = other.Negate
}

// Implementors provides automatic JSON unmarshalling with a list of struct implementing (directly or indirectly) ComparisonBasic
func (bcb *BaseComparisonBasic) Implementors() []ComparisonBasic {
	return []ComparisonBasic{
		new(CustomApplicationTypeComparison),
		new(MobilePlatformComparison),
		new(ApplicationTypeComparison),
		new(BitnessComparision),
		new(PaasTypeComparison),
		new(OSArchitectureComparison),
		new(ServiceTopologyComparison),
		new(StringComparison),
		new(DatabaseTopologyComparison),
		new(DCRumDecoderComparison),
		new(IndexedTagComparison),
		new(TagComparison),
		new(HypervisorTypeComparision),
		new(CloudTypeComparison),
		new(IndexedNameComparison),
		new(IndexedStringComparison),
		new(SimpleTechComparison),
		new(AzureSkuComparision),
		new(EntityIDComparison),
		new(SimpleHostTechComparison),
		new(IntegerComparison),
		new(ServiceTypeComparison),
		new(OSTypeComparison),
		new(SyntheticEngineTypeComparison),
		new(IPAddressComparison),
		new(AzureComputeModeComparison),
	}
}

// UnmarshalJSON provides custom JSON deserialization
func (bcb *BaseComparisonBasic) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, bcb)
}

// MarshalJSON provides custom JSON serialization
func (bcb *BaseComparisonBasic) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(bcb)
}
