package requestattributes

// Normalization String values transformation.
//  If the **dataType** is not `string`, set the `Original` here.
type Normalization string

// Normalizations offers the known enum values
var Normalizations = struct {
	Original    Normalization
	ToLowerCase Normalization
	ToUpperCase Normalization
}{
	"ORIGINAL",
	"TO_LOWER_CASE",
	"TO_UPPER_CASE",
}
