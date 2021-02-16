package autotags

// CustomHostMetadataKeySource The source of the custom metadata.
type CustomHostMetadataKeySource string

// CustomHostMetadataKeySources offers the known enum values
var CustomHostMetadataKeySources = struct {
	Environment         CustomHostMetadataKeySource
	GoogleComputeEngine CustomHostMetadataKeySource
	Plugin              CustomHostMetadataKeySource
}{
	"ENVIRONMENT",
	"GOOGLE_COMPUTE_ENGINE",
	"PLUGIN",
}
