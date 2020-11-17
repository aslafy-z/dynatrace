package requestattributes

// IibNodeType The IBM integration bus node type for which the value is captured.
//  This or `iibMethodNodeCondition` is required if the **source** is: `IIB_NODE`.
//  Not applicable in other cases.
type IibNodeType string

// IibNodeTypes offers the known enum values
var IibNodeTypes = struct {
	AggregateControlNode       IibNodeType
	AggregateReplyNode         IibNodeType
	AggregateRequestNode       IibNodeType
	CallableFlowReplyNode      IibNodeType
	CollectorNode              IibNodeType
	ComputeNode                IibNodeType
	DatabaseNode               IibNodeType
	DecisionServiceNode        IibNodeType
	DotNetComputeNode          IibNodeType
	FileReadNode               IibNodeType
	FilterNode                 IibNodeType
	FlowOrderNode              IibNodeType
	GroupCompleteNode          IibNodeType
	GroupGatherNode            IibNodeType
	GroupScatterNode           IibNodeType
	HTTPHeader                 IibNodeType
	JavaComputeNode            IibNodeType
	JmsClientReceive           IibNodeType
	JmsClientReplyNode         IibNodeType
	JmsHeader                  IibNodeType
	MqGetNode                  IibNodeType
	MqOutputNode               IibNodeType
	PassthruNode               IibNodeType
	ResetContentDescriptorNode IibNodeType
	ReSequenceNode             IibNodeType
	RouteNode                  IibNodeType
	SAPReplyNode               IibNodeType
	ScaReplyNode               IibNodeType
	SecurityPep                IibNodeType
	SequenceNode               IibNodeType
	SoapExtractNode            IibNodeType
	SoapReplyNode              IibNodeType
	SoapWrapperNode            IibNodeType
	SrRetrieveEntityNode       IibNodeType
	SrRetrieveItServiceNode    IibNodeType
	ThrowNode                  IibNodeType
	TraceNode                  IibNodeType
	TryCatchNode               IibNodeType
	ValidateNode               IibNodeType
	WsReplyNode                IibNodeType
	XslMqsiNode                IibNodeType
}{
	"AGGREGATE_CONTROL_NODE",
	"AGGREGATE_REPLY_NODE",
	"AGGREGATE_REQUEST_NODE",
	"CALLABLE_FLOW_REPLY_NODE",
	"COLLECTOR_NODE",
	"COMPUTE_NODE",
	"DATABASE_NODE",
	"DECISION_SERVICE_NODE",
	"DOT_NET_COMPUTE_NODE",
	"FILE_READ_NODE",
	"FILTER_NODE",
	"FLOW_ORDER_NODE",
	"GROUP_COMPLETE_NODE",
	"GROUP_GATHER_NODE",
	"GROUP_SCATTER_NODE",
	"HTTP_HEADER",
	"JAVA_COMPUTE_NODE",
	"JMS_CLIENT_RECEIVE",
	"JMS_CLIENT_REPLY_NODE",
	"JMS_HEADER",
	"MQ_GET_NODE",
	"MQ_OUTPUT_NODE",
	"PASSTHRU_NODE",
	"RESET_CONTENT_DESCRIPTOR_NODE",
	"RE_SEQUENCE_NODE",
	"ROUTE_NODE",
	"SAP_REPLY_NODE",
	"SCA_REPLY_NODE",
	"SECURITY_PEP",
	"SEQUENCE_NODE",
	"SOAP_EXTRACT_NODE",
	"SOAP_REPLY_NODE",
	"SOAP_WRAPPER_NODE",
	"SR_RETRIEVE_ENTITY_NODE",
	"SR_RETRIEVE_IT_SERVICE_NODE",
	"THROW_NODE",
	"TRACE_NODE",
	"TRY_CATCH_NODE",
	"VALIDATE_NODE",
	"WS_REPLY_NODE",
	"XSL_MQSI_NODE",
}
