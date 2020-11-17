package customservices

import "encoding/json"

// FileNameMatcher has no documentation
type FileNameMatcher string

// FileNameMatchers offers the known enum values
var FileNameMatchers = struct {
	EndsWith   FileNameMatcher
	Equals     FileNameMatcher
	StartsWith FileNameMatcher
}{
	"ENDS_WITH",
	"EQUALS",
	"STARTS_WITH",
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *FileNameMatcher) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = FileNameMatcher(name)
	return nil
}
