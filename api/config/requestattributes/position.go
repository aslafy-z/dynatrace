package requestattributes

// Position The position of the extracted string relative to delimiters.
type Position string

// Positions offers the known enum values
var Positions = struct {
	After   Position
	Before  Position
	Between Position
}{
	"AFTER",
	"BEFORE",
	"BETWEEN",
}
