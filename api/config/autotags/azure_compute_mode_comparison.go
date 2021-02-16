package autotags

import "github.com/dtcookie/gojson"

// AzureComputeModeComparison Comparison for `AZURE_COMPUTE_MODE` attributes.
type AzureComputeModeComparison struct {
	BaseComparisonBasic `json:",type=AZURE_COMPUTE_MODE"`
	Operator            AzureComputeModeComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *AzureComputeModeComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (acmc *AzureComputeModeComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, acmc)
}

// MarshalJSON provides custom JSON serialization
func (acmc *AzureComputeModeComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(acmc)
}
