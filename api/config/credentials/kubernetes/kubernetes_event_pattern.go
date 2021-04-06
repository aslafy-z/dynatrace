package kubernetes

import "github.com/dtcookie/gojson"

// KubernetesEventPattern Represents a single Kubernetes events field selector (=event filter based on the K8s field selector).
type KubernetesEventPattern struct {
	Active        bool   `json:"active"`        // Whether subscription to this events field selector is enabled (value set to `true`). If disabled (value set to `false`), Dynatrace will stop fetching events from the Kubernetes API for this events field selector
	FieldSelector string `json:"fieldSelector"` // The field selector string (url decoding is applied) when storing it.
	Label         string `json:"label"`         // A label of the events field selector.
}

// UnmarshalJSON provides custom JSON deserialization
func (kep *KubernetesEventPattern) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, kep)
}

// MarshalJSON provides custom JSON serialization
func (kep *KubernetesEventPattern) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(kep)
}
