package autotags

import "github.com/dtcookie/gojson"

// IPAddressComparison Comparison for `IP_ADDRESS` attributes.
type IPAddressComparison struct {
	BaseComparisonBasic `json:",type=IP_ADDRESS"`
	CaseSensitive       *bool                       `json:"caseSensitive,omitempty"` // The comparison is case-sensitive (`true`) or insensitive (`false`).
	Operator            IPAddressComparisonOperator `json:"operator"`                // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *string                     `json:"value,omitempty"`         // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (iac *IPAddressComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, iac)
}

// MarshalJSON provides custom JSON serialization
func (iac *IPAddressComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(iac)
}
