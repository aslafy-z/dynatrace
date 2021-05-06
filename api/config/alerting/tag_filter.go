package alerting

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// TagFilter A tag-based filter of monitored entities.
type TagFilter struct {
	Context  Context                    `json:"context"`         // The origin of the tag, such as AWS or Cloud Foundry.  Custom tags use the `CONTEXTLESS` value.
	Key      string                     `json:"key"`             // The key of the tag.  Custom tags have the tag value here.
	Value    *string                    `json:"value,omitempty"` // The value of the tag.   Not applicable to custom tags.
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *TagFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"context": {
			Type:        hcl.TypeString,
			Description: "The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the `CONTEXTLESS` value",
			Required:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "The key of the tag. Custom tags have the tag value here",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value of the tag. Not applicable to custom tags",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *TagFilter) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["context"] = string(me.Context)
	result["key"] = me.Key
	if me.Value != nil {
		result["value"] = *me.Value
	}
	return result, nil
}

func (me *TagFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "context")
		delete(me.Unknowns, "key")
		delete(me.Unknowns, "value")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("context"); ok {
		me.Context = Context(value.(string))
	}
	if value, ok := decoder.GetOk("key"); ok {
		me.Key = value.(string)
	}
	if value, ok := decoder.GetOk("value"); ok {
		me.Value = opt.NewString(value.(string))
	}
	return nil
}

func (me *TagFilter) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(me.Context)
		if err != nil {
			return nil, err
		}
		m["context"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(me.Key)
		if err != nil {
			return nil, err
		}
		m["key"] = rawMessage
	}
	if me.Value != nil {
		rawMessage, err := json.Marshal(me.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}

	return json.Marshal(m)
}

func (me *TagFilter) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["context"]; found {
		if err := json.Unmarshal(v, &me.Context); err != nil {
			return err
		}
	}
	if v, found := m["key"]; found {
		if err := json.Unmarshal(v, &me.Key); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &me.Value); err != nil {
			return err
		}
	}

	delete(m, "context")
	delete(m, "key")
	delete(m, "value")

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// Context The origin of the tag, such as AWS or Cloud Foundry.
// Custom tags use the `CONTEXTLESS` value.
type Context string

// Contexts offers the known enum values
var Contexts = struct {
	AWS          Context
	AWSGeneric   Context
	Azure        Context
	CloudFoundry Context
	Contextless  Context
	Environment  Context
	GoogleCloud  Context
	Kubernetes   Context
}{
	"AWS",
	"AWS_GENERIC",
	"AZURE",
	"CLOUD_FOUNDRY",
	"CONTEXTLESS",
	"ENVIRONMENT",
	"GOOGLE_CLOUD",
	"KUBERNETES",
}
