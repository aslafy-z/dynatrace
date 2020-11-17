package managementzones

import "github.com/dtcookie/gojson"

// ConditionKey The key to identify the data we're matching.
// The actual set of fields and possible values vary, depending on the **type** of the key.
// Find the list of actual objects in the description of the **type** field.
type ConditionKey interface {
	Initialize(*BaseConditionKey) // Initialize duplicates the properties of the other AbstractConditionKey into this instance. It also serves as a unique method for structs implementing the interface this abstract class is derived from
	Implementors() []ConditionKey // Implementors provides automatic JSON unmarshalling with a list of struct implementing (directly or indirectly) ConditionKey
}

// BaseConditionKey The key to identify the data we're matching.
// The actual set of fields and possible values vary, depending on the **type** of the key.
// Find the list of actual objects in the description of the **type** field.
type BaseConditionKey struct {
	Attribute Attribute         `json:"attribute"`      // The attribute to be used for comparision.
	Type      *ConditionKeyType `json:"type,omitempty"` // Defines the actual set of fields depending on the value. See one of the following objects:  * `PROCESS_CUSTOM_METADATA_KEY` -> CustomProcessMetadataConditionKey  * `HOST_CUSTOM_METADATA_KEY` -> CustomHostMetadataConditionKey  * `PROCESS_PREDEFINED_METADATA_KEY` -> ProcessMetadataConditionKey  * `STRING` -> StringConditionKey
}

// Initialize duplicates the properties of the other BaseConditionKey into this instance
// It also serves as a unique method for structs implementing the interface this abstract class is derived from
func (bck *BaseConditionKey) Initialize(other *BaseConditionKey) {
	bck.Attribute = other.Attribute
	bck.Type = other.Type
}

// Implementors provides automatic JSON unmarshalling with a list of struct implementing (directly or indirectly) ConditionKey
func (bck *BaseConditionKey) Implementors() []ConditionKey {
	return []ConditionKey{
		new(CustomHostMetadataConditionKey),
		new(StringConditionKey),
		new(CustomProcessMetadataConditionKey),
		new(ProcessMetadataConditionKey),
	}
}

// UnmarshalJSON provides custom JSON deserialization
func (bck *BaseConditionKey) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, bck)
}

// MarshalJSON provides custom JSON serialization
func (bck *BaseConditionKey) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(bck)
}
