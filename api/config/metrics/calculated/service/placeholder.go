package service

import "github.com/dtcookie/dynatrace/api/config/metrics/calculated/service/propagation"

// Placeholder The custom placeholder to be used as a naming value pattern.
//  It enables you to extract a request attribute value or other request attribute and use it in the naming pattern.
type Placeholder struct {
	Aggregation       *Aggregation        `json:"aggregation,omitempty"`       // Which value of the request attribute must be used when it occurs across multiple child requests.  Only applicable for the `SERVICE_REQUEST_ATTRIBUTE` attribute, when **useFromChildCalls** is `true`.  For the `COUNT` aggregation, the **kind** field is not applicable.
	Attribute         Attribute           `json:"attribute"`                   // The attribute to extract from. You can only use attributes of the **string** type.
	Source            *propagation.Source `json:"source,omitempty"`            // Defines valid sources of request attributes for conditions or placeholders.
	Normalization     *Normalization      `json:"normalization,omitempty"`     // The format of the extracted string.
	DelimiterOrRegex  *string             `json:"delimiterOrRegex,omitempty"`  // Depending on the **type** value:   * `REGEX_EXTRACTION`: The regular expression.   * `BETWEEN_DELIMITER`: The opening delimiter string to look for.   * All other values: The delimiter string to look for.
	UseFromChildCalls *bool               `json:"useFromChildCalls,omitempty"` // If `true` request attribute will be taken from a child service call.   Only applicable for the `SERVICE_REQUEST_ATTRIBUTE` attribute. Defaults to `false`.
	RequestAttribute  *string             `json:"requestAttribute,omitempty"`  // The request attribute to extract from.   Required if the **kind** value is `SERVICE_REQUEST_ATTRIBUTE`. Not applicable otherwise.
	EndDelimiter      *string             `json:"endDelimiter,omitempty"`      // The closing delimiter string to look for.   Required if the **kind** value is `BETWEEN_DELIMITER`. Not applicable otherwise.
	Kind              Kind                `json:"kind"`                        // The type of extraction.   Defines either usage of regular expression (`regex`) or the position of request attribute value to be extracted.  When the **attribute** is `SERVICE_REQUEST_ATTRIBUTE` attribute and **aggregation** is `COUNT`, needs to be set to `ORIGINAL_TEXT`
	Name              string              `json:"name"`                        // The name of the placeholder. Use it in the naming pattern as `{name}`.
}

// Aggregation Which value of the request attribute must be used when it occurs across multiple child requests.
// Only applicable for the `SERVICE_REQUEST_ATTRIBUTE` attribute, when **useFromChildCalls** is `true`.
// For the `COUNT` aggregation, the **kind** field is not applicable.
type Aggregation string

// Aggregations offers the known enum values
var Aggregations = struct {
	Count Aggregation
	First Aggregation
	Last  Aggregation
}{
	"COUNT",
	"FIRST",
	"LAST",
}

// Normalization The format of the extracted string.
type Normalization string

// Normalizations offers the known enum values
var Normalizations = struct {
	Original    Normalization
	ToLowerCase Normalization
	ToUpperCase Normalization
}{
	"ORIGINAL",
	"TO_LOWER_CASE",
	"TO_UPPER_CASE",
}

// Kind The type of extraction.
//  Defines either usage of regular expression (`regex`) or the position of request attribute value to be extracted.
// When the **attribute** is `SERVICE_REQUEST_ATTRIBUTE` attribute and **aggregation** is `COUNT`, needs to be set to `ORIGINAL_TEXT`
type Kind string

// Kinds offers the known enum values
var Kinds = struct {
	AfterDelimiter   Kind
	BeforeDelimiter  Kind
	BetweenDelimiter Kind
	OriginalText     Kind
	RegexExtraction  Kind
}{
	"AFTER_DELIMITER",
	"BEFORE_DELIMITER",
	"BETWEEN_DELIMITER",
	"ORIGINAL_TEXT",
	"REGEX_EXTRACTION",
}
