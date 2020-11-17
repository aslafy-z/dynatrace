package requestattributes

// FileNameMatcher The operator of the comparison.
//  If not set, `EQUALS` is used.
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
