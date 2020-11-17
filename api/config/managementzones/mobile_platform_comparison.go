package managementzones

import "github.com/dtcookie/gojson"

// MobilePlatformComparison Comparison for `MOBILE_PLATFORM` attributes.
type MobilePlatformComparison struct {
	BaseComparisonBasic `json:",type=MOBILE_PLATFORM"`
	Operator            MobilePlatformComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *MobilePlatformComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (mpc *MobilePlatformComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, mpc)
}

// MarshalJSON provides custom JSON serialization
func (mpc *MobilePlatformComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(mpc)
}
