package autotags

// CustomProcessMetadataKeySource The source of the custom metadata.
type CustomProcessMetadataKeySource string

// CustomProcessMetadataKeySources offers the known enum values
var CustomProcessMetadataKeySources = struct {
	CloudFoundry CustomProcessMetadataKeySource
	Environment  CustomProcessMetadataKeySource
	GoogleCloud  CustomProcessMetadataKeySource
	Kubernetes   CustomProcessMetadataKeySource
	Plugin       CustomProcessMetadataKeySource
}{
	"CLOUD_FOUNDRY",
	"ENVIRONMENT",
	"GOOGLE_CLOUD",
	"KUBERNETES",
	"PLUGIN",
}
