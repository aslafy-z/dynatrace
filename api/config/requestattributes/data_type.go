package requestattributes

// DataType The data type of the request attribute.
type DataType string

// DataTypes offers the known enum values
var DataTypes = struct {
	Double  DataType
	Integer DataType
	String  DataType
}{
	"DOUBLE",
	"INTEGER",
	"STRING",
}
