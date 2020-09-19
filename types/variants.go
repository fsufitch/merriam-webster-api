package types

// WithVariants is a compositing type for parsing the `vrs` property
type WithVariants struct {
	// https://dictionaryapi.com/products/json#sec-2.vrs
	Variants []Variant `json:"vrs"`
}

// Variant https://dictionaryapi.com/products/json#sec-2.vrs
type Variant struct {
	Text  string `json:"va"`
	Label string `json:"vl"`
	WithPronounciations
	WithSenseSpecificInflectionPluralLabel
}
