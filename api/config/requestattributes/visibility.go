package requestattributes

// Visibility The visibility of the method to capture.
type Visibility string

// Visibilitys offers the known enum values
var Visibilitys = struct {
	Internal         Visibility
	PackageProtected Visibility
	Private          Visibility
	Protected        Visibility
	Public           Visibility
}{
	"INTERNAL",
	"PACKAGE_PROTECTED",
	"PRIVATE",
	"PROTECTED",
	"PUBLIC",
}
