package managementzones

import "github.com/dtcookie/gojson"

// DCRumDecoderComparison Comparison for `DCRUM_DECODER_TYPE` attributes.
type DCRumDecoderComparison struct {
	BaseComparisonBasic `json:",type=DCRUM_DECODER_TYPE"`
	Operator            DCRumDecoderComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *DCRumDecoderComparisonValue   `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (ddc *DCRumDecoderComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ddc)
}

// MarshalJSON provides custom JSON serialization
func (ddc *DCRumDecoderComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ddc)
}
