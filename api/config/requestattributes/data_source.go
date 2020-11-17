package requestattributes

import "github.com/dtcookie/gojson"

// DataSource has no documentation
type DataSource struct {
	CapturingAndStorageLocation *CapturingAndStorageLocation `json:"capturingAndStorageLocation,omitempty"` // Specifies the location where the values are captured and stored.  Required if the **source** is one of the following: `GET_PARAMETER`, `URI`, `REQUEST_HEADER`, `RESPONSE_HEADER`.   Not applicable in other cases.   If the **source** value is `REQUEST_HEADER` or `RESPONSE_HEADER`, the `CAPTURE_AND_STORE_ON_BOTH` location is not allowed.
	Scope                       *ScopeConditions             `json:"scope,omitempty"`                       // Conditions for data capturing.
	ParameterName               *string                      `json:"parameterName,omitempty"`               // The name of the web request parameter to capture.  Required if the **source** is one of the following: `POST_PARAMETER`, `GET_PARAMETER`, `REQUEST_HEADER`, `RESPONSE_HEADER`, `CUSTOM_ATTRIBUTE`.  Not applicable in other cases.
	IibMethodNodeCondition      *ValueCondition              `json:"iibMethodNodeCondition,omitempty"`      // IBM integration bus label node name condition for which the value is captured.
	Methods                     []CapturedMethod             `json:"methods,omitempty"`                     // The method specification if the **source** value is `METHOD_PARAM`.   Not applicable in other cases.
	SessionAttributeTechnology  *SessionAttributeTechnology  `json:"sessionAttributeTechnology,omitempty"`  // The technology of the session attribute to capture if the **source** value is `SESSION_ATTRIBUTE`. \n\n Not applicable in other cases.
	Technology                  *Technology                  `json:"technology,omitempty"`                  // The technology of the method to capture if the **source** value is `METHOD_PARAM`. \n\n Not applicable in other cases.
	ValueProcessing             *ValueProcessing             `json:"valueProcessing,omitempty"`             // Process values as specified.
	CICSSDKMethodNodeCondition  *ValueCondition              `json:"cicsSDKMethodNodeCondition,omitempty"`  // IBM integration bus label node name condition for which the value is captured.
	Enabled                     bool                         `json:"enabled"`                               // The data source is enabled (`true`) or disabled (`false`).
	Source                      Source                       `json:"source"`                                // The source of the attribute to capture. Works in conjunction with **parameterName** or **methods** and **technology**.
	IibLabelMethodNodeCondition *ValueCondition              `json:"iibLabelMethodNodeCondition,omitempty"` // IBM integration bus label node name condition for which the value is captured.
	IibNodeType                 *IibNodeType                 `json:"iibNodeType,omitempty"`                 // The IBM integration bus node type for which the value is captured.  This or `iibMethodNodeCondition` is required if the **source** is: `IIB_NODE`.  Not applicable in other cases.
}

// UnmarshalJSON provides custom JSON deserialization
func (ds *DataSource) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ds)
}

// MarshalJSON provides custom JSON serialization
func (ds *DataSource) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ds)
}
