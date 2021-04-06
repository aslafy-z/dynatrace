package aws

// ParameterLocation has no documentation
type ParameterLocation string

// ParameterLocations offers the known enum values
var ParameterLocations = struct {
	Path        ParameterLocation
	PayloadBody ParameterLocation
	Query       ParameterLocation
}{
	"PATH",
	"PAYLOAD_BODY",
	"QUERY",
}
