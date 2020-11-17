package customservices

import "encoding/json"

// Technology has no documentation
type Technology string

// Technologies offers the known enum values
var Technologies = struct {
	DotNet Technology
	Go     Technology
	Java   Technology
	NodeJS Technology
	PHP    Technology
}{
	"dotNet",
	"go",
	"java",
	"nodeJS",
	"php",
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *Technology) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = Technology(name)
	return nil
}
