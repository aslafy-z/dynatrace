package managementzones

import "github.com/dtcookie/gojson"

// CustomHostMetadataConditionKey The key for dynamic attributes of the `HOST_CUSTOM_METADATA_KEY` type.
type CustomHostMetadataConditionKey struct {
	BaseConditionKey `json:",type=HOST_CUSTOM_METADATA_KEY"`
	DynamicKey       CustomHostMetadataKey `json:"dynamicKey"` // The key of the attribute, which need dynamic keys.  Not applicable otherwise, as the attibute itself acts as a key.
}

// UnmarshalJSON provides custom JSON deserialization
func (chmck *CustomHostMetadataConditionKey) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, chmck)
}

// MarshalJSON provides custom JSON serialization
func (chmck *CustomHostMetadataConditionKey) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(chmck)
}
