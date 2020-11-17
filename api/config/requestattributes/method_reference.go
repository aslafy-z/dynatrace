package requestattributes

import "github.com/dtcookie/gojson"

// MethodReference Configuration of a method to be captured.
type MethodReference struct {
	ReturnType      string           `json:"returnType"`                // The return type.
	Visibility      Visibility       `json:"visibility"`                // The visibility of the method to capture.
	ArgumentTypes   []string         `json:"argumentTypes"`             // The list of argument types.
	ClassName       *string          `json:"className,omitempty"`       // The class name where the method to capture resides.   Either this or the **fileName** must be set.
	FileName        *string          `json:"fileName,omitempty"`        // The file name where the method to capture resides.   Either this or **className** must be set.
	FileNameMatcher *FileNameMatcher `json:"fileNameMatcher,omitempty"` // The operator of the comparison.   If not set, `EQUALS` is used.
	MethodName      string           `json:"methodName"`                // The name of the method to capture.
	Modifiers       []Modifier       `json:"modifiers"`                 // The modifiers of the method to capture.
}

// UnmarshalJSON provides custom JSON deserialization
func (mr *MethodReference) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, mr)
}

// MarshalJSON provides custom JSON serialization
func (mr *MethodReference) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(mr)
}
