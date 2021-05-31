package comparisoninfo

// ComparisonInfo Type-specific comparison for attributes. The actual set of fields depends on the `type` of the comparison.
// See the [Service metrics API - JSON models](https://dt-url.net/9803svb) help topic for example models of every notification type.
type ComparisonInfo interface {
	GetType() Type
}

// BaseComparisonInfo Type-specific comparison for attributes. The actual set of fields depends on the `type` of the comparison.
// See the [Service metrics API - JSON models](https://dt-url.net/9803svb) help topic for example models of every notification type.
type BaseComparisonInfo struct {
	Negate bool `json:"negate"` // Reverse the comparison **operator**. For example, it turns **equals** into **does not equal**.
	Type   Type `json:"type"`   // Defines the actual set of fields depending on the value. See one of the following objects:  * `STRING` -> StringComparisonInfo  * `NUMBER` -> NumberComparisonInfo  * `BOOLEAN` -> BooleanComparisonInfo  * `HTTP_METHOD` -> HttpMethodComparisonInfo  * `STRING_REQUEST_ATTRIBUTE` -> StringRequestAttributeComparisonInfo  * `NUMBER_REQUEST_ATTRIBUTE` -> NumberRequestAttributeComparisonInfo  * `ZOS_CALL_TYPE` -> ZosComparisonInfo  * `IIB_INPUT_NODE_TYPE` -> IIBInputNodeTypeComparisonInfo  * `ESB_INPUT_NODE_TYPE` -> ESBInputNodeTypeComparisonInfo  * `FAILED_STATE` -> FailedStateComparisonInfo  * `FLAW_STATE` -> FlawStateComparisonInfo  * `FAILURE_REASON` -> FailureReasonComparisonInfo  * `HTTP_STATUS_CLASS` -> HttpStatusClassComparisonInfo  * `TAG` -> TagComparisonInfo  * `FAST_STRING` -> FastStringComparisonInfo  * `SERVICE_TYPE` -> ServiceTypeComparisonInfo
}

func (me *BaseComparisonInfo) GetType() Type {
	return me.Type
}
