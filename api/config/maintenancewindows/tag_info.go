package maintenancewindows

// TagInfo Tag of a Dynatrace entity.
type TagInfo struct {
	Context Context `json:"context"`         // The origin of the tag, such as AWS or Cloud Foundry.   Custom tags use the `CONTEXTLESS` value.
	Key     string  `json:"key"`             // The key of the tag.   Custom tags have the tag value here.
	Value   *string `json:"value,omitempty"` // The value of the tag.   Not applicable to custom tags.
}
