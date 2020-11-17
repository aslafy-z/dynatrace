package managementzones

import "github.com/dtcookie/gojson"

// CustomProcessMetadataConditionKey The key for dynamic attributes of the `PROCESS_CUSTOM_METADATA_KEY` type.
type CustomProcessMetadataConditionKey struct {
	BaseConditionKey `json:",type=PROCESS_CUSTOM_METADATA_KEY"`
	DynamicKey       CustomProcessMetadataKey `json:"dynamicKey"` // The key of the attribute, which need dynamic keys.  Not applicable otherwise, as the attibute itself acts as a key.
}

// UnmarshalJSON provides custom JSON deserialization
func (cpmck *CustomProcessMetadataConditionKey) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, cpmck)
}

// MarshalJSON provides custom JSON serialization
func (cpmck *CustomProcessMetadataConditionKey) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(cpmck)
}
