package comparisoninfo

// HTTPMethod Comparison for `HTTP_METHOD` attributes.
type HTTPMethod struct {
	BaseComparisonInfo                      // `json:",type=HTTP_METHOD"`
	Comparison         HTTPMethodComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *HTTPMethodValue     `json:"value,omitempty"`  // The value to compare to.
	Values             []HTTPMethodValue    `json:"values,omitempty"` // The values to compare to.
}

// HTTPMethodComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type HTTPMethodComparison string

// HTTPMethodComparisons offers the known enum values
var HTTPMethodComparisons = struct {
	Equals      HTTPMethodComparison
	EqualsAnyOf HTTPMethodComparison
	Exists      HTTPMethodComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// HTTPMethodValue The values to compare to.
type HTTPMethodValue string

// HTTPMethodValues offers the known enum values
var HTTPMethodValues = struct {
	Connect HTTPMethodValue
	Delete  HTTPMethodValue
	Get     HTTPMethodValue
	Head    HTTPMethodValue
	Options HTTPMethodValue
	Patch   HTTPMethodValue
	Post    HTTPMethodValue
	Put     HTTPMethodValue
	Trace   HTTPMethodValue
}{
	"CONNECT",
	"DELETE",
	"GET",
	"HEAD",
	"OPTIONS",
	"PATCH",
	"POST",
	"PUT",
	"TRACE",
}
