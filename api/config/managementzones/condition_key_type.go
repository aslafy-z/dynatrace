package managementzones

// ConditionKeyType Defines the actual set of fields depending on the value. See one of the following objects:
// * `PROCESS_CUSTOM_METADATA_KEY` -> CustomProcessMetadataConditionKey
// * `HOST_CUSTOM_METADATA_KEY` -> CustomHostMetadataConditionKey
// * `PROCESS_PREDEFINED_METADATA_KEY` -> ProcessMetadataConditionKey
// * `STRING` -> StringConditionKey
type ConditionKeyType string

// ConditionKeyTypes offers the known enum values
var ConditionKeyTypes = struct {
	HostCustomMetadataKey        ConditionKeyType
	ProcessCustomMetadataKey     ConditionKeyType
	ProcessPredefinedMetadataKey ConditionKeyType
	String                       ConditionKeyType
}{
	"HOST_CUSTOM_METADATA_KEY",
	"PROCESS_CUSTOM_METADATA_KEY",
	"PROCESS_PREDEFINED_METADATA_KEY",
	"STRING",
}
