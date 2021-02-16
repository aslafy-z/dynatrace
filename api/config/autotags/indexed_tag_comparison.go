package autotags

import "github.com/dtcookie/gojson"

// IndexedTagComparison Comparison for `INDEXED_TAG` attributes.
type IndexedTagComparison struct {
	BaseComparisonBasic `json:",type=INDEXED_TAG"`
	Operator            IndexedTagComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *TagInfo                     `json:"value,omitempty"` // Tag of a Dynatrace entity.
}

// UnmarshalJSON provides custom JSON deserialization
func (itc *IndexedTagComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, itc)
}

// MarshalJSON provides custom JSON serialization
func (itc *IndexedTagComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(itc)
}
