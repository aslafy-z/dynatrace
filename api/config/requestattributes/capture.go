package requestattributes

// Capture What to capture from the method.
type Capture string

// Captures offers the known enum values
var Captures = struct {
	Argument        Capture
	ClassName       Capture
	MethodName      Capture
	Occurrences     Capture
	SimpleClassName Capture
	This            Capture
}{
	"ARGUMENT",
	"CLASS_NAME",
	"METHOD_NAME",
	"OCCURRENCES",
	"SIMPLE_CLASS_NAME",
	"THIS",
}
