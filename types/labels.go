package types

// WithFunctionalLabel is a compositing type for parsing the `fl` property
type WithFunctionalLabel struct {
	// https://dictionaryapi.com/products/json#sec-2.fl
	Function string `json:"fl,omitempty"`
}

// WithGeneralLabels is a compositing type for parsing the `lbs` property
type WithGeneralLabels struct {
	// https://dictionaryapi.com/products/json#sec-2.lbs
	Labels []string `json:"lbs,omitempty"`
}

// WithSubjectStatusLabels is a compositing type for parsing the `sls` property
type WithSubjectStatusLabels struct {
	// https://dictionaryapi.com/products/json#sec-2.sls
	SubjectStatusLabels []string `json:"sls,omitempty"`
}

// WithParenthesizedSubjectStatusLabel is a compositing type for parsing the `psl` property
type WithParenthesizedSubjectStatusLabel struct {
	// https://dictionaryapi.com/products/json#sec-2.psl
	ParenthesizedSubjectStatusLabel string `json:"psl,omitempty"`
}

// WithSenseSpecificInflectionPluralLabel is a compositing type for parsing the `spl` property
type WithSenseSpecificInflectionPluralLabel struct {
	// https://dictionaryapi.com/products/json#sec-2.spl
	SenseSpecificInflectionPluralLabel string `json:"spl,omitempty"`
}

// WithSenseSpecificGrammaticalLabel is a compositing type for parsing the `sgram` property
type WithSenseSpecificGrammaticalLabel struct {
	// https://dictionaryapi.com/products/json#sec-2.sgram
	SenseSpecificGrammaticalLabel string `json:"sgram,omitempty"`
}
