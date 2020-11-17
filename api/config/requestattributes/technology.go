package requestattributes

// Technology The technology of the method to capture if the **source** value is `METHOD_PARAM`. \n\n Not applicable in other cases.
type Technology string

// Technologys offers the known enum values
var Technologys = struct {
	DotNet Technology
	Java   Technology
	PHP    Technology
}{
	"DOTNET",
	"JAVA",
	"PHP",
}
