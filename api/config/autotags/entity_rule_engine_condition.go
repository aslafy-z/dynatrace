package autotags

import "github.com/dtcookie/gojson"

// EntityRuleEngineCondition A condition defines how to execute matching logic for an entity.
type EntityRuleEngineCondition struct {
	ComparisonInfo ComparisonBasic `json:"comparisonInfo"` // Defines how the matching is actually performed: what and how are we comparing.  The actual set of fields and possible values of the **operator** field depend on the **type** of the comparison. \n\nFind the list of actual models in the description of the **type** field and check the description of the model you need.
	Key            ConditionKey    `json:"key"`            // The key to identify the data we're matching.  The actual set of fields and possible values vary, depending on the **type** of the key.  Find the list of actual objects in the description of the **type** field.
}

// UnmarshalJSON provides custom JSON deserialization
func (erec *EntityRuleEngineCondition) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, erec)
}

// MarshalJSON provides custom JSON serialization
func (erec *EntityRuleEngineCondition) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(erec)
}
