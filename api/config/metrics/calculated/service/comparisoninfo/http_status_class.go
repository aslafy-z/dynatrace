package comparisoninfo

// HTTPStatusClass Comparison for `HTTP_STATUS_CLASS` attributes.
type HTTPStatusClass struct {
	BaseComparisonInfo                           // `json:",type=HTTP_STATUS_CLASS"`
	Values             []HTTPStatusClassValue    `json:"values,omitempty"` // The values to compare to.
	Comparison         HTTPStatusClassComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value              *HTTPStatusClassValue     `json:"value,omitempty"`  // The value to compare to.
}

// HTTPStatusClassComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type HTTPStatusClassComparison string

// HTTPStatusClassComparisons offers the known enum values
var HTTPStatusClassComparisons = struct {
	Equals      HTTPStatusClassComparison
	EqualsAnyOf HTTPStatusClassComparison
	Exists      HTTPStatusClassComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// HTTPStatusClassValue The value to compare to.
type HTTPStatusClassValue string

// HTTPStatusClassValues offers the known enum values
var HTTPStatusClassValues = struct {
	C1xx       HTTPStatusClassValue
	C2xx       HTTPStatusClassValue
	C3xx       HTTPStatusClassValue
	C4xx       HTTPStatusClassValue
	C5xx       HTTPStatusClassValue
	NoResponse HTTPStatusClassValue
}{
	"C_1XX",
	"C_2XX",
	"C_3XX",
	"C_4XX",
	"C_5XX",
	"NO_RESPONSE",
}
