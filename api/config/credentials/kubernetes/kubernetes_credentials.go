package kubernetes

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// KubernetesCredentials Configuration for specific Kubernetes credentials.
type KubernetesCredentials struct {
	EventsIntegrationEnabled   *bool                      `json:"eventsIntegrationEnabled,omitempty"`   // The monitoring of events is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. Event monitoring depends on the active state of this configuration to be true.  If not set on creation, the `false` value is used.  If the field is omitted during an update, the old value remains unaffected.
	Active                     *bool                      `json:"active,omitempty"`                     // The monitoring is enabled (`true`) or disabled (`false`) for given credentials configuration.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	EndpointStatusInfo         *string                    `json:"endpointStatusInfo,omitempty"`         // The detailed status info of the configured endpoint.
	WorkloadIntegrationEnabled *bool                      `json:"workloadIntegrationEnabled,omitempty"` // Workload and cloud application processing is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected.
	EndpointStatus             *EndpointStatus            `json:"endpointStatus,omitempty"`             // The status of the configured endpoint. ASSIGNED: The credentials are assigned to an ActiveGate which is responsible for processing. UNASSIGNED: The credentials are not yet assigned to an ActiveGate so there is currently no processing. DISABLED: The credentials have been disabled by the user. FASTCHECK_AUTH_ERROR: The credentials are invalid. FASTCHECK_TLS_ERROR: The endpoint TLS certificate is invalid. FASTCHECK_NO_RESPONSE: The endpoint did not return a result until the timeout was reached. FASTCHECK_INVALID_ENDPOINT: The endpoint URL was invalid. FASTCHECK_AUTH_LOCKED: The credentials seem to be locked. UNKNOWN: An unknown error occured.
	Label                      string                     `json:"label"`                                // The name of the Kubernetes credentials configuration.  Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`. Leading or trailing whitespace is not allowed.
	ID                         *string                    `json:"id,omitempty"`                         // The ID of the given credentials configuration.
	Metadata                   *api.ConfigurationMetadata `json:"metadata,omitempty"`                   // Metadata useful for debugging
	AuthToken                  *string                    `json:"authToken,omitempty"`                  // The service account bearer token for the Kubernetes API server.  Submit your token on creation or update of the configuration. For security reasons, GET requests return this field as `null`.  If the field is omitted during an update, the old value remains unaffected.
	CertificateCheckEnabled    *bool                      `json:"certificateCheckEnabled,omitempty"`    // The check of SSL certificates is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	EndpointURL                string                     `json:"endpointUrl"`                          // The URL of the Kubernetes API server.  It must be unique within a Dynatrace environment.  The URL must valid according to RFC 2396. Leading or trailing whitespaces are not allowed.
	EventsFieldSelectors       []*KubernetesEventPattern  `json:"eventsFieldSelectors,omitempty"`       // Kubernetes event filters based on field-selectors. If set to `null` on creation, no events field selectors are subscribed. If set to `null` on update, no change of stored events field selectors is applied. Set an empty list to clear all events field selectors.
	Unknowns                   map[string]json.RawMessage `json:"-"`
}

func (kc *KubernetesCredentials) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
		"endpoint_url": {
			Type:        hcl.TypeString,
			Description: "The URL of the Kubernetes API server.  It must be unique within a Dynatrace environment.  The URL must valid according to RFC 2396. Leading or trailing whitespaces are not allowed.",
			Optional:    true,
		},
		"label": {
			Type:        hcl.TypeString,
			Description: "The name of the Kubernetes credentials configuration.  Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`. Leading or trailing whitespace is not allowed.",
			Required:    true,
		},
		"events_integration_enabled": {
			Type:        hcl.TypeBool,
			Description: "Monitoring of events is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. Event monitoring depends on the active state of this configuration to be true.  If not set on creation, the `false` value is used.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"active": {
			Type:        hcl.TypeBool,
			Description: "Monitoring is enabled (`true`) or disabled (`false`) for given credentials configuration.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"workload_integration_enabled": {
			Type:        hcl.TypeBool,
			Description: "Workload and cloud application processing is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"auth_token": {
			Type:        hcl.TypeString,
			Description: "The service account bearer token for the Kubernetes API server.  Submit your token on creation or update of the configuration. For security reasons, GET requests return this field as `null`.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"certificate_check_enabled": {
			Type:        hcl.TypeBool,
			Description: "The check of SSL certificates is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
			Default:     true,
		},
		"events_field_selectors": {
			Type:        hcl.TypeList,
			Description: "The check of SSL certificates is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(KubernetesEventPattern).Schema(),
			},
		},
		/*
			"metadata": {
				Type:        hcl.TypeString,
				Description: "Workload and cloud application processing is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected.",
				Optional:    true,
			},
			"endpoint_status": {
				Type:        hcl.TypeString,
				Description: "Workload and cloud application processing is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected.",
				Optional:    true,
			},
			"id": {
				Type:        hcl.TypeString,
				Description: "Workload and cloud application processing is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected.",
				Optional:    true,
			},
		*/
	}
}

func (kc *KubernetesCredentials) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if kc.Unknowns != nil {
		data, err := json.Marshal(kc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if kc.EventsIntegrationEnabled != nil {
		result["events_integration_enabled"] = *kc.EventsIntegrationEnabled
	}
	if kc.Active != nil {
		result["active"] = *kc.Active
	}
	if kc.WorkloadIntegrationEnabled != nil {
		result["workload_integration_enabled"] = *kc.WorkloadIntegrationEnabled
	}
	result["label"] = kc.Label
	if kc.AuthToken != nil {
		result["auth_token"] = *kc.AuthToken
	}
	if kc.CertificateCheckEnabled != nil {
		result["certificate_check_enabled"] = *kc.CertificateCheckEnabled
	}
	result["endpoint_url"] = kc.EndpointURL
	if kc.EventsFieldSelectors != nil {
		entries := []interface{}{}
		for _, eventPattern := range kc.EventsFieldSelectors {
			if marshalled, err := eventPattern.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["events_field_selectors"] = entries
	}
	return result, nil
}

func (kc *KubernetesCredentials) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), kc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &kc.Unknowns); err != nil {
			return err
		}
		delete(kc.Unknowns, "eventsIntegrationEnabled")
		delete(kc.Unknowns, "active")
		delete(kc.Unknowns, "endpointStatusInfo")
		delete(kc.Unknowns, "workloadIntegrationEnabled")
		delete(kc.Unknowns, "endpointStatus")
		delete(kc.Unknowns, "label")
		delete(kc.Unknowns, "id")
		delete(kc.Unknowns, "metadata")
		delete(kc.Unknowns, "authToken")
		delete(kc.Unknowns, "certificateCheckEnabled")
		delete(kc.Unknowns, "endpointUrl")
		delete(kc.Unknowns, "eventsFieldSelectors")
	}
	if _, value := decoder.GetChange("events_integration_enabled"); value != nil {
		kc.EventsIntegrationEnabled = opt.NewBool(value.(bool))
	}
	if _, value := decoder.GetChange("active"); value != nil {
		kc.Active = opt.NewBool(value.(bool))
	}
	if _, value := decoder.GetChange("workload_integration_enabled"); value != nil {
		kc.WorkloadIntegrationEnabled = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("label"); ok {
		kc.Label = value.(string)
	}
	if value, ok := decoder.GetOk("id"); ok {
		kc.ID = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("auth_token"); ok {
		kc.AuthToken = opt.NewString(value.(string))
	}
	if _, value := decoder.GetChange("certificate_check_enabled"); value != nil {
		kc.CertificateCheckEnabled = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("endpoint_url"); ok {
		kc.EndpointURL = value.(string)
	}
	if result, ok := decoder.GetOk("events_field_selectors.#"); ok {
		kc.EventsFieldSelectors = []*KubernetesEventPattern{}
		for idx := 0; idx < result.(int); idx++ {
			eventPattern := new(KubernetesEventPattern)
			if err := eventPattern.UnmarshalHCL(hcl.NewDecoder(decoder, "events_field_selectors", idx)); err != nil {
				return err
			}
			kc.EventsFieldSelectors = append(kc.EventsFieldSelectors, eventPattern)
		}
	}
	return nil
}

// UnmarshalJSON provides custom JSON deserialization
func (kc *KubernetesCredentials) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["eventsIntegrationEnabled"]; found {
		if err := json.Unmarshal(v, &kc.EventsIntegrationEnabled); err != nil {
			return err
		}
	}
	if v, found := m["active"]; found {
		if err := json.Unmarshal(v, &kc.Active); err != nil {
			return err
		}
	}
	if v, found := m["endpointStatusInfo"]; found {
		if err := json.Unmarshal(v, &kc.EndpointStatusInfo); err != nil {
			return err
		}
	}
	if v, found := m["workloadIntegrationEnabled"]; found {
		if err := json.Unmarshal(v, &kc.WorkloadIntegrationEnabled); err != nil {
			return err
		}
	}
	if v, found := m["endpointStatus"]; found {
		if err := json.Unmarshal(v, &kc.EndpointStatus); err != nil {
			return err
		}
	}
	if v, found := m["label"]; found {
		if err := json.Unmarshal(v, &kc.Label); err != nil {
			return err
		}
	}
	if v, found := m["id"]; found {
		if err := json.Unmarshal(v, &kc.ID); err != nil {
			return err
		}
	}
	if v, found := m["metadata"]; found {
		if err := json.Unmarshal(v, &kc.Metadata); err != nil {
			return err
		}
	}
	if v, found := m["authToken"]; found {
		if err := json.Unmarshal(v, &kc.AuthToken); err != nil {
			return err
		}
	}
	if v, found := m["certificateCheckEnabled"]; found {
		if err := json.Unmarshal(v, &kc.CertificateCheckEnabled); err != nil {
			return err
		}
	}
	if v, found := m["endpointUrl"]; found {
		if err := json.Unmarshal(v, &kc.EndpointURL); err != nil {
			return err
		}
	}
	if v, found := m["eventsFieldSelectors"]; found {
		if err := json.Unmarshal(v, &kc.EventsFieldSelectors); err != nil {
			return err
		}
	}
	return nil
}

func (kc *KubernetesCredentials) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if kc.Unknowns != nil {
		for k, v := range kc.Unknowns {
			m[k] = v
		}
	}
	if kc.EventsIntegrationEnabled != nil {
		rawMessage, err := json.Marshal(kc.EventsIntegrationEnabled)
		if err != nil {
			return nil, err
		}
		m["eventsIntegrationEnabled"] = rawMessage
	}
	if kc.Active != nil {
		rawMessage, err := json.Marshal(kc.Active)
		if err != nil {
			return nil, err
		}
		m["active"] = rawMessage
	}
	if kc.EndpointStatusInfo != nil {
		rawMessage, err := json.Marshal(kc.EndpointStatusInfo)
		if err != nil {
			return nil, err
		}
		m["endpointStatusInfo"] = rawMessage
	}
	if kc.WorkloadIntegrationEnabled != nil {
		rawMessage, err := json.Marshal(kc.WorkloadIntegrationEnabled)
		if err != nil {
			return nil, err
		}
		m["workloadIntegrationEnabled"] = rawMessage
	}
	if kc.EndpointStatus != nil {
		rawMessage, err := json.Marshal(kc.EndpointStatus)
		if err != nil {
			return nil, err
		}
		m["endpointStatus"] = rawMessage
	}
	rawMessage, err := json.Marshal(kc.Label)
	if err != nil {
		return nil, err
	}
	m["label"] = rawMessage
	if kc.ID != nil {
		rawMessage, err := json.Marshal(kc.ID)
		if err != nil {
			return nil, err
		}
		m["id"] = rawMessage
	}
	if kc.Metadata != nil {
		rawMessage, err := json.Marshal(kc.Metadata)
		if err != nil {
			return nil, err
		}
		m["metadata"] = rawMessage
	}
	if kc.AuthToken != nil {
		rawMessage, err := json.Marshal(kc.AuthToken)
		if err != nil {
			return nil, err
		}
		m["authToken"] = rawMessage
	}
	if kc.CertificateCheckEnabled != nil {
		rawMessage, err := json.Marshal(kc.CertificateCheckEnabled)
		if err != nil {
			return nil, err
		}
		m["certificateCheckEnabled"] = rawMessage
	}
	rawMessage, err = json.Marshal(kc.EndpointURL)
	if err != nil {
		return nil, err
	}
	m["endpointUrl"] = rawMessage
	if kc.EventsFieldSelectors != nil {
		rawMessage, err := json.Marshal(kc.EventsFieldSelectors)
		if err != nil {
			return nil, err
		}
		m["eventsFieldSelectors"] = rawMessage
	}
	return json.Marshal(m)
}
