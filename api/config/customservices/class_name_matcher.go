package customservices

import "encoding/json"

// ClassNameMatcher has no documentation
type ClassNameMatcher string

// ClassNameMatchers offers the known enum values
var ClassNameMatchers = struct {
	EndsWith   ClassNameMatcher
	Equals     ClassNameMatcher
	StartsWith ClassNameMatcher
}{
	"ENDS_WITH",
	"EQUALS",
	"STARTS_WITH",
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *ClassNameMatcher) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = ClassNameMatcher(name)
	return nil
}
