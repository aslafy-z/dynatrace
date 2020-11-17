package managementzones

// CustomApplicationTypeComparisonValue The value to compare to.
type CustomApplicationTypeComparisonValue string

// CustomApplicationTypeComparisonValues offers the known enum values
var CustomApplicationTypeComparisonValues = struct {
	AmazonEcho        CustomApplicationTypeComparisonValue
	Desktop           CustomApplicationTypeComparisonValue
	Embedded          CustomApplicationTypeComparisonValue
	Iot               CustomApplicationTypeComparisonValue
	MicrosoftHololens CustomApplicationTypeComparisonValue
	Ufo               CustomApplicationTypeComparisonValue
}{
	"AMAZON_ECHO",
	"DESKTOP",
	"EMBEDDED",
	"IOT",
	"MICROSOFT_HOLOLENS",
	"UFO",
}
