package comparisoninfo

import (
	"fmt"

	"github.com/dtcookie/hcl"
)

type Wrapper struct {
	Negate     bool
	Comparison ComparisonInfo
}

func (me *Wrapper) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"negate": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "Reverse the comparison **operator**. For example, it turns **equals** into **does not equal**",
		},
		"boolean": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Boolean Comparison for `BOOLEAN` attributes",
			Elem:          &hcl.Resource{Schema: new(Boolean).Schema()},
		},
		"esb_input_node_type": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Type-specific comparison information for attributes of type 'ESB_INPUT_NODE_TYPE'",
			Elem:          &hcl.Resource{Schema: new(ESBInputNodeType).Schema()},
		},
		"failed_state": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `FAILED_STATE` attributes",
			Elem:          &hcl.Resource{Schema: new(FailedState).Schema()},
		},
		"failure_reason": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `FAILURE_REASON` attributes",
			Elem:          &hcl.Resource{Schema: new(FailureReason).Schema()},
		},
		"fast_string": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `FAST_STRING` attributes. Use it for all service property attributes",
			Elem:          &hcl.Resource{Schema: new(FastString).Schema()},
		},
		"flaw_state": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `FLAW_STATE` attributes",
			Elem:          &hcl.Resource{Schema: new(FlawState).Schema()},
		},
		"http_method": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `HTTP_METHOD` attributes",
			Elem:          &hcl.Resource{Schema: new(HTTPMethod).Schema()},
		},
		"http_status_class": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `HTTP_STATUS_CLASS` attributes",
			Elem:          &hcl.Resource{Schema: new(HTTPStatusClass).Schema()},
		},
		"iib_input_node_type": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `IIB_INPUT_NODE_TYPE` attributes",
			Elem:          &hcl.Resource{Schema: new(IIBInputNodeType).Schema()},
		},
		"number": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `NUMBER` attributes",
			Elem:          &hcl.Resource{Schema: new(Number).Schema()},
		},
		"number_request_attribute": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "service_type", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `NUMBER_REQUEST_ATTRIBUTE` attributes",
			Elem:          &hcl.Resource{Schema: new(NumberRequestAttribute).Schema()},
		},
		"service_type": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "string", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `SERVICE_TYPE` attributes",
			Elem:          &hcl.Resource{Schema: new(ServiceType).Schema()},
		},
		"string": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string_request_attribute", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `STRING` attributes",
			Elem:          &hcl.Resource{Schema: new(String).Schema()},
		},
		"string_request_attribute": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "tag", "zos_call_type", "generic"},
			Description:   "Comparison for `STRING_REQUEST_ATTRIBUTE` attributes",
			Elem:          &hcl.Resource{Schema: new(StringRequestAttribute).Schema()},
		},
		"tag": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "zos_call_type", "generic"},
			Description:   "Comparison for `TAG` attributes",
			Elem:          &hcl.Resource{Schema: new(Tag).Schema()},
		},
		"zos_call_type": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "generic"},
			Description:   "Comparison for `ZOS_CALL_TYPE` attributes",
			Elem:          &hcl.Resource{Schema: new(ZOSCallType).Schema()},
		},
		"generic": {
			Type:          hcl.TypeList,
			Optional:      true,
			MaxItems:      1,
			MinItems:      1,
			ConflictsWith: []string{"boolean", "esb_input_node_type", "failed_state", "failure_reason", "fast_string", "flaw_state", "http_method", "http_status_class", "iib_input_node_type", "number", "number_request_attribute", "service_type", "string", "string_request_attribute", "tag", "zos_call_type"},
			Description:   "Comparison for `NUMBER` attributes",
			Elem:          &hcl.Resource{Schema: new(BaseComparisonInfo).Schema()},
		},
	}
}

func (me *Wrapper) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	properties.Encode("negate", me.Negate)
	switch cmp := me.Comparison.(type) {
	case *Boolean:
		if err := properties.Encode("boolean", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *ESBInputNodeType:
		if err := properties.Encode("esb_input_node_type", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *FailedState:
		if err := properties.Encode("failed_state", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *FailureReason:
		if err := properties.Encode("failure_reason", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *FastString:
		if err := properties.Encode("fast_string", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *FlawState:
		if err := properties.Encode("flaw_state", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *HTTPMethod:
		if err := properties.Encode("http_method", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *HTTPStatusClass:
		if err := properties.Encode("http_status_class", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *IIBInputNodeType:
		if err := properties.Encode("iib_input_node_type", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *NumberRequestAttribute:
		if err := properties.Encode("number_request_attribute", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *Number:
		if err := properties.Encode("number", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *ServiceType:
		if err := properties.Encode("service_type", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *StringRequestAttribute:
		if err := properties.Encode("string_request_attribute", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *String:
		if err := properties.Encode("string", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *Tag:
		if err := properties.Encode("tag", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *ZOSCallType:
		if err := properties.Encode("zos_call_type", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *BaseComparisonInfo:
		if err := properties.Encode("generic", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	default:
		return nil, fmt.Errorf("cannot HCL marshal objects of type %T", cmp)
	}
}

func (me *Wrapper) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("negate", &me.Negate); err != nil {
		return err
	}
	var err error
	var cmp interface{}
	if cmp, err = decoder.DecodeAny(map[string]interface{}{
		"boolean":                  new(Boolean),
		"esb_input_node_type":      new(ESBInputNodeType),
		"failed_state":             new(FailedState),
		"failure_reason":           new(FailureReason),
		"fast_string":              new(FastString),
		"flaw_state":               new(FlawState),
		"http_method":              new(HTTPMethod),
		"http_status_class":        new(HTTPStatusClass),
		"iib_input_node_type":      new(IIBInputNodeType),
		"number":                   new(Number),
		"number_request_attribute": new(NumberRequestAttribute),
		"service_type":             new(ServiceType),
		"string":                   new(String),
		"string_request_attribute": new(StringRequestAttribute),
		"tag":                      new(Tag),
		"zos_call_type":            new(ZOSCallType),
		"generic":                  new(BaseComparisonInfo)}); err != nil {
		return err
	}
	if cmp != nil {
		me.Comparison = cmp.(ComparisonInfo)
	}
	return nil
}
