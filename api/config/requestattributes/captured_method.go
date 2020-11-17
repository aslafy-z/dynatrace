package requestattributes

import "github.com/dtcookie/gojson"

// CapturedMethod has no documentation
type CapturedMethod struct {
	ArgumentIndex    *int32          `json:"argumentIndex,omitempty"`    // The index of the argument to capture. Set `0` to capture the return value, `1` or higher to capture a mehtod argument.   Required if the **capture** is set to `ARGUMENT`.  Not applicable in other cases.
	Capture          Capture         `json:"capture"`                    // What to capture from the method.
	DeepObjectAccess *string         `json:"deepObjectAccess,omitempty"` // The getter chain to apply to the captured object. It is required in one of the following cases:  The **capture** is set to `THIS`.    The **capture** is set to `ARGUMENT`, and the argument is not a primitive, a primitive wrapper class, a string, or an array.   Not applicable in other cases.
	Method           MethodReference `json:"method"`                     // Configuration of a method to be captured.
}

// UnmarshalJSON provides custom JSON deserialization
func (cm *CapturedMethod) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, cm)
}

// MarshalJSON provides custom JSON serialization
func (cm *CapturedMethod) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(cm)
}
