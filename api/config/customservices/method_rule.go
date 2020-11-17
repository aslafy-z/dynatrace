package customservices

// MethodRule TODO: documentation
type MethodRule struct {
	ID            string   `json:"id,omitempty"`            // The ID of the method rule
	MethodName    string   `json:"methodName"`              // The method to instrument
	ArgumentTypes []string `json:"argumentTypes,omitempty"` // Fully qualified types of argument the method expects
	ReturnType    string   `json:"returnType"`              // Fully qualified type the method returns
}
