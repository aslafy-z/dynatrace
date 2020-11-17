package requestattributes

// SessionAttributeTechnology The technology of the session attribute to capture if the **source** value is `SESSION_ATTRIBUTE`. \n\n Not applicable in other cases.
type SessionAttributeTechnology string

// SessionAttributeTechnologys offers the known enum values
var SessionAttributeTechnologys = struct {
	ASPNet     SessionAttributeTechnology
	ASPNetCore SessionAttributeTechnology
	Java       SessionAttributeTechnology
}{
	"ASP_NET",
	"ASP_NET_CORE",
	"JAVA",
}
