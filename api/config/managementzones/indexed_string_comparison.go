package managementzones

import "github.com/dtcookie/gojson"

// IndexedStringComparison Comparison for `INDEXED_STRING` attributes.
type IndexedStringComparison struct {
	BaseComparisonBasic `json:",type=INDEXED_STRING"`
	Operator            IndexedStringComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *string                         `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (isc *IndexedStringComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, isc)
}

// MarshalJSON provides custom JSON serialization
func (isc *IndexedStringComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(isc)
}
