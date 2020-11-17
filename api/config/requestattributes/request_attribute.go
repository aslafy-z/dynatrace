package requestattributes

import "github.com/dtcookie/gojson"

// RequestAttribute has no documentation
type RequestAttribute struct {
	SkipPersonalDataMasking bool                   `json:"skipPersonalDataMasking"` // Personal data masking flag. Set `true` to skip masking.   Warning: This will potentially access personalized data.
	Confidential            bool                   `json:"confidential"`            // Confidential data flag. Set `true` to treat the captured data as confidential.
	DataSources             []DataSource           `json:"dataSources"`             // The list of data sources.
	DataType                DataType               `json:"dataType"`                // The data type of the request attribute.
	ID                      *string                `json:"id,omitempty"`            // The ID of the request attribute.
	Metadata                *ConfigurationMetadata `json:"metadata,omitempty"`      // Metadata useful for debugging
	Normalization           Normalization          `json:"normalization"`           // String values transformation.   If the **dataType** is not `string`, set the `Original` here.
	Enabled                 bool                   `json:"enabled"`                 // The request attribute is enabled (`true`) or disabled (`false`).
	Name                    string                 `json:"name"`                    // The name of the request attribute.
	Aggregation             Aggregation            `json:"aggregation"`             // Aggregation type for the request values.
}

// UnmarshalJSON provides custom JSON deserialization
func (ra *RequestAttribute) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, ra)
}

// MarshalJSON provides custom JSON serialization
func (ra *RequestAttribute) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(ra)
}
