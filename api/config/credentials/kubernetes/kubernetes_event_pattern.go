package kubernetes

import (
	"encoding/json"

	"github.com/dtcookie/gojson"
	"github.com/dtcookie/hcl"
)

// KubernetesEventPattern Represents a single Kubernetes events field selector (=event filter based on the K8s field selector).
type KubernetesEventPattern struct {
	Active        bool                       `json:"active"`        // Whether subscription to this events field selector is enabled (value set to `true`). If disabled (value set to `false`), Dynatrace will stop fetching events from the Kubernetes API for this events field selector
	FieldSelector string                     `json:"fieldSelector"` // The field selector string (url decoding is applied) when storing it.
	Label         string                     `json:"label"`         // A label of the events field selector.
	Unknowns      map[string]json.RawMessage `json:"-"`
}

func (kep *KubernetesEventPattern) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
		"active": {
			Type:        hcl.TypeBool,
			Description: "Whether subscription to this events field selector is enabled (value set to `true`). If disabled (value set to `false`), Dynatrace will stop fetching events from the Kubernetes API for this events field selector",
			Required:    true,
		},
		"field_selector": {
			Type:        hcl.TypeString,
			Description: "The field selector string (url decoding is applied) when storing it.",
			Required:    true,
		},
		"label": {
			Type:        hcl.TypeString,
			Description: "A label of the events field selector.",
			Required:    true,
		},
	}
}

func (kep *KubernetesEventPattern) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if kep.Unknowns != nil {
		data, err := json.Marshal(kep.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}

	result["active"] = kep.Active
	result["field_selector"] = kep.FieldSelector
	result["label"] = kep.Label

	return result, nil
}

func (kep *KubernetesEventPattern) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), kep); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &kep.Unknowns); err != nil {
			return err
		}
		delete(kep.Unknowns, "active")
		delete(kep.Unknowns, "fieldSelector")
		delete(kep.Unknowns, "label")
	}
	if value, ok := decoder.GetOk("active"); ok {
		kep.Active = value.(bool)
	}
	if value, ok := decoder.GetOk("field_selector"); ok {
		kep.FieldSelector = value.(string)
	}
	if value, ok := decoder.GetOk("label"); ok {
		kep.Label = value.(string)
	}
	return nil
}

// UnmarshalJSON provides custom JSON deserialization
func (kep *KubernetesEventPattern) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, kep)
}

// MarshalJSON provides custom JSON serialization
func (kep *KubernetesEventPattern) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(kep)
}
