package alertingprofiles

// IncludeMode The filtering mode:
// * `INCLUDE_ANY`: The rule applies to monitored entities that have at least one of the specified tags. You can specify up to 100 tags.
// * `INCLUDE_ALL`: The rule applies to monitored entities that have **all** of the specified tags. You can specify up to 10 tags.
// * `NONE`: The rule applies to all monitored entities.
type IncludeMode string

// IncludeModes offers the known enum values
var IncludeModes = struct {
	IncludeAll IncludeMode
	IncludeAny IncludeMode
	None       IncludeMode
}{
	"INCLUDE_ALL",
	"INCLUDE_ANY",
	"NONE",
}
