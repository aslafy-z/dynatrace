package managementzones

// IPAddressComparisonOperator Operator of the comparison. You can reverse it by setting **negate** to `true`.
// Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
type IPAddressComparisonOperator string

// IPAddressComparisonOperators offers the known enum values
var IPAddressComparisonOperators = struct {
	BeginsWith   IPAddressComparisonOperator
	Contains     IPAddressComparisonOperator
	EndsWith     IPAddressComparisonOperator
	Equals       IPAddressComparisonOperator
	Exists       IPAddressComparisonOperator
	IsIPInRange  IPAddressComparisonOperator
	RegexMatches IPAddressComparisonOperator
}{
	"BEGINS_WITH",
	"CONTAINS",
	"ENDS_WITH",
	"EQUALS",
	"EXISTS",
	"IS_IP_IN_RANGE",
	"REGEX_MATCHES",
}
