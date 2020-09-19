package types

// WithDividedSense is a compositing type for parsing the `sdsense` property
type WithDividedSense struct {
	DividedSense *DividedSense `json:"sdsense"`
}

// DividedSense https://dictionaryapi.com/products/json#sec-2.sdsense
type DividedSense struct {
	Divider string `json:"sd"`
	WithEtymology
	WithInflections
	WithGeneralLabels
	WithPronounciations
	WithSenseSpecificGrammaticalLabel
	WithSubjectStatusLabels
	WithDefiningText
}
