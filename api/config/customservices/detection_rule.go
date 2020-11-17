package customservices

// DetectionRule the defining rules for a CustomService
type DetectionRule struct {
	ID               string           `json:"id,omitempty"`              // The ID of the detection rule
	Enabled          bool             `json:"enabled"`                   // Rule enabled/disabled
	FileName         string           `json:"fileName,omitempty"`        // The PHP file containing the class or methods to instrument. Required for PHP custom service. Not applicable to Java and .NET
	FileNameMatcher  FileNameMatcher  `json:"fileNameMatcher,omitempty"` // Matcher applying to the file name. Default value is `ENDS_WITH` (if applicable)
	ClassName        string           `json:"className,omitempty"`       // The fully qualified class or interface to instrument. Required for Java and .NET custom services. Not applicable to PHP
	ClassNameMatcher ClassNameMatcher `json:"matcher,omitempty"`         // Matcher applying to the class name. `STARTS_WITH` can only be used if there is at least one annotation defined. Default value is `EQUALS`
	MethodRules      []MethodRule     `json:"methodRules"`               // List of methods to instrument
	Annotations      []string         `json:"annotations,omitempty"`     // Additional annotations filter of the rule. Only classes where all listed annotations are available in the class itself or any of its superclasses are instrumented. nNot applicable to PHP
}
