package comparisoninfo

// TagInfo Tag of a Dynatrace entity.
type TagInfo struct {
	Context TagInfoContext `json:"context"`         // The origin of the tag, such as AWS or Cloud Foundry.   Custom tags use the `CONTEXTLESS` value.
	Key     string         `json:"key"`             // The key of the tag.   Custom tags have the tag value here.
	Value   *string        `json:"value,omitempty"` // The value of the tag.   Not applicable to custom tags.
}

// TagInfoContext The origin of the tag, such as AWS or Cloud Foundry.
//  Custom tags use the `CONTEXTLESS` value.
type TagInfoContext string

// TagInfoContexts offers the known enum values
var TagInfoContexts = struct {
	AWS          TagInfoContext
	AWSGeneric   TagInfoContext
	Azure        TagInfoContext
	CloudFoundry TagInfoContext
	Contextless  TagInfoContext
	Environment  TagInfoContext
	GoogleCloud  TagInfoContext
	Kubernetes   TagInfoContext
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
