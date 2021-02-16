package autotags

import "github.com/dtcookie/gojson"

// ProcessMetadataConditionKey The key for dynamic attributes of the `PROCESS_PREDEFINED_METADATA_KEY` type.
type ProcessMetadataConditionKey struct {
	BaseConditionKey `json:",type=PROCESS_PREDEFINED_METADATA_KEY"`
	DynamicKey       DynamicKey `json:"dynamicKey"` // The key of the attribute, which need dynamic keys.  Not applicable otherwise, as the attibute itself acts as a key.
}

// UnmarshalJSON provides custom JSON deserialization
func (pmck *ProcessMetadataConditionKey) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, pmck)
}

// MarshalJSON provides custom JSON serialization
func (pmck *ProcessMetadataConditionKey) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(pmck)
}
