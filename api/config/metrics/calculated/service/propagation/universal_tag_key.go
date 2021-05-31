package propagation

// UniversalTagKey has no documentation
type UniversalTagKey struct {
	Key     *string                 `json:"key,omitempty"`     // has no documentation
	Context *UniversalTagKeyContext `json:"context,omitempty"` // has no documentation
}

// UniversalTagKeyContext has no documentation
type UniversalTagKeyContext string

// UniversalTagKeyContexts offers the known enum values
var UniversalTagKeyContexts = struct {
	AWS                 UniversalTagKeyContext
	AWSGeneric          UniversalTagKeyContext
	Azure               UniversalTagKeyContext
	CloudFoundry        UniversalTagKeyContext
	Contextless         UniversalTagKeyContext
	Environment         UniversalTagKeyContext
	GoogleComputeEngine UniversalTagKeyContext
	Kubernetes          UniversalTagKeyContext
}{
	"AWS",
	"AWS_GENERIC",
	"AZURE",
	"CLOUD_FOUNDRY",
	"CONTEXTLESS",
	"ENVIRONMENT",
	"GOOGLE_COMPUTE_ENGINE",
	"KUBERNETES",
}
