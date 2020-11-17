package maintenancewindows

// TagCombination The logic that applies when several tags are specified: AND/OR.
// If not set, the OR logic is used.
type TagCombination string

// TagCombinations offers the known enum values
var TagCombinations = struct {
	And TagCombination
	Or  TagCombination
}{
	"AND",
	"OR",
}
