package propagation

// Source Defines valid sources of request attributes for conditions or placeholders.
type Source struct {
	ManagementZone *string       `json:"managementZone,omitempty"` // Use only request attributes from services that belong to this management zone.. Use either this or `serviceTag`
	ServiceTag     *UniversalTag `json:"serviceTag,omitempty"`     // Use only request attributes from services that have this tag. Use either this or `managementZone`
}
