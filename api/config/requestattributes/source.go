package requestattributes

// Source The source of the attribute to capture. Works in conjunction with **parameterName** or **methods** and **technology**.
type Source string

// Sources offers the known enum values
var Sources = struct {
	CICSSdk          Source
	ClientIP         Source
	CustomAttribute  Source
	IibLabel         Source
	IibNode          Source
	MethodParam      Source
	PostParameter    Source
	QueryParameter   Source
	RequestHeader    Source
	ResponseHeader   Source
	SessionAttribute Source
	URI              Source
	URIPath          Source
}{
	"CICS_SDK",
	"CLIENT_IP",
	"CUSTOM_ATTRIBUTE",
	"IIB_LABEL",
	"IIB_NODE",
	"METHOD_PARAM",
	"POST_PARAMETER",
	"QUERY_PARAMETER",
	"REQUEST_HEADER",
	"RESPONSE_HEADER",
	"SESSION_ATTRIBUTE",
	"URI",
	"URI_PATH",
}
