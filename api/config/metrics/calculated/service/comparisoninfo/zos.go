package comparisoninfo

// Zos Comparison for `ZOS_CALL_TYPE` attributes.
type ZOS struct {
	BaseComparisonInfo               //`json:",type=ZOS_CALL_TYPE"`
	Comparison         ZOSComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *ZosValue     `json:"value,omitempty"`  // The value to compare to.
	Values             []ZosValue    `json:"values,omitempty"` // The values to compare to.
}

// ZOSComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type ZOSComparison string

// ZOSComparisons offers the known enum values
var ZOSComparisons = struct {
	Equals      ZOSComparison
	EqualsAnyOf ZOSComparison
	Exists      ZOSComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// ZosValue The value to compare to.
type ZosValue string

// ZosValues offers the known enum values
var ZosValues = struct {
	CTG             ZosValue
	Dpl             ZosValue
	ExplicitAdk     ZosValue
	IMSConnect      ZosValue
	IMSConnectApi   ZosValue
	IMSItra         ZosValue
	IMSMsc          ZosValue
	IMSPgmSwitch    ZosValue
	IMSSharedQueues ZosValue
	IMSTransExec    ZosValue
	Mq              ZosValue
	Soap            ZosValue
	Start           ZosValue
	Tx              ZosValue
	Unknown         ZosValue
}{
	"CTG",
	"DPL",
	"EXPLICIT_ADK",
	"IMS_CONNECT",
	"IMS_CONNECT_API",
	"IMS_ITRA",
	"IMS_MSC",
	"IMS_PGM_SWITCH",
	"IMS_SHARED_QUEUES",
	"IMS_TRANS_EXEC",
	"MQ",
	"SOAP",
	"START",
	"TX",
	"UNKNOWN",
}
