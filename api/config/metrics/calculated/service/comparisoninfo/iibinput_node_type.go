package comparisoninfo

// IIBInputNodeType Comparison for `IIB_INPUT_NODE_TYPE` attributes.
type IIBInputNodeType struct {
	BaseComparisonInfo                            // `json:",type=IIB_INPUT_NODE_TYPE"`
	Value              *IIBInputNodeTypeValue     `json:"value,omitempty"`  // The value to compare to.
	Values             []IIBInputNodeTypeValue    `json:"values,omitempty"` // The values to compare to.
	Comparison         IIBInputNodeTypeComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
}

// IIBInputNodeTypeComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type IIBInputNodeTypeComparison string

// IIBInputNodeTypeComparisons offers the known enum values
var IIBInputNodeTypeComparisons = struct {
	Equals      IIBInputNodeTypeComparison
	EqualsAnyOf IIBInputNodeTypeComparison
	Exists      IIBInputNodeTypeComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// IIBInputNodeTypeValue The value to compare to.
type IIBInputNodeTypeValue string

// IIBInputNodeTypeValues offers the known enum values
var IIBInputNodeTypeValues = struct {
	CallableFlowAsyncResponseNode IIBInputNodeTypeValue
	CallableFlowInputNode         IIBInputNodeTypeValue
	DatabaseInputNode             IIBInputNodeTypeValue
	DotNetInputNode               IIBInputNodeTypeValue
	EmailInputNode                IIBInputNodeTypeValue
	EventInput                    IIBInputNodeTypeValue
	EventInputNode                IIBInputNodeTypeValue
	FileInputNode                 IIBInputNodeTypeValue
	FteInputNode                  IIBInputNodeTypeValue
	HTTPAsyncResponse             IIBInputNodeTypeValue
	JdEdwardsInputNode            IIBInputNodeTypeValue
	JmsClientInputNode            IIBInputNodeTypeValue
	LabelNode                     IIBInputNodeTypeValue
	MqInputNode                   IIBInputNodeTypeValue
	PeopleSoftInputNode           IIBInputNodeTypeValue
	RestAsyncResponse             IIBInputNodeTypeValue
	RestRequest                   IIBInputNodeTypeValue
	SAPInputNode                  IIBInputNodeTypeValue
	ScaAsyncResponseNode          IIBInputNodeTypeValue
	ScaInputNode                  IIBInputNodeTypeValue
	SiebelInputNode               IIBInputNodeTypeValue
	SoapInputNode                 IIBInputNodeTypeValue
	TcpipClientInputNode          IIBInputNodeTypeValue
	TcpipClientRequestNode        IIBInputNodeTypeValue
	TcpipServerInputNode          IIBInputNodeTypeValue
	TcpipServerRequestNode        IIBInputNodeTypeValue
	TimeoutNotificationNode       IIBInputNodeTypeValue
	WsInputNode                   IIBInputNodeTypeValue
}{
	"CALLABLE_FLOW_ASYNC_RESPONSE_NODE",
	"CALLABLE_FLOW_INPUT_NODE",
	"DATABASE_INPUT_NODE",
	"DOTNET_INPUT_NODE",
	"EMAIL_INPUT_NODE",
	"EVENT_INPUT",
	"EVENT_INPUT_NODE",
	"FILE_INPUT_NODE",
	"FTE_INPUT_NODE",
	"HTTP_ASYNC_RESPONSE",
	"JD_EDWARDS_INPUT_NODE",
	"JMS_CLIENT_INPUT_NODE",
	"LABEL_NODE",
	"MQ_INPUT_NODE",
	"PEOPLE_SOFT_INPUT_NODE",
	"REST_ASYNC_RESPONSE",
	"REST_REQUEST",
	"SAP_INPUT_NODE",
	"SCA_ASYNC_RESPONSE_NODE",
	"SCA_INPUT_NODE",
	"SIEBEL_INPUT_NODE",
	"SOAP_INPUT_NODE",
	"TCPIP_CLIENT_INPUT_NODE",
	"TCPIP_CLIENT_REQUEST_NODE",
	"TCPIP_SERVER_INPUT_NODE",
	"TCPIP_SERVER_REQUEST_NODE",
	"TIMEOUT_NOTIFICATION_NODE",
	"WS_INPUT_NODE",
}
