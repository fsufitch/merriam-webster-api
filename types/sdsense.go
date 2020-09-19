package types

// WithDividedSense is a compositing type for parsing the `sdsense` property
type WithDividedSense struct {
	DividedSense *DividedSense `json:"sdsense,omitempty"`
}

// DividedSense https://dictionaryapi.com/products/json#sec-2.sdsense
type DividedSense struct {
	Divider string `json:"sd,omitempty"`
	WithEtymology
	WithInflections
	WithGeneralLabels
	WithPronounciations
	WithSenseSpecificGrammaticalLabel
	WithSubjectStatusLabels
	WithDefiningText
}
