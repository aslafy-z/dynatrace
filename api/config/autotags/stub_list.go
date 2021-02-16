package autotags

import "github.com/dtcookie/gojson"

// StubList An ordered list of short representations of Dynatrace entities.
type StubList struct {
	Values []EntityShortRepresentation `json:"values"` // An ordered list of short representations of Dynatrace entities.
}

// UnmarshalJSON provides custom JSON deserialization
func (sl *StubList) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, sl)
}

// MarshalJSON provides custom JSON serialization
func (sl *StubList) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(sl)
}
