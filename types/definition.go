package types

// WithDefinitions is a compositing type for parsing the `def` property
type WithDefinitions struct {
	Definitions []Definition `json:"def,omitempty"`
}

// Definition https://dictionaryapi.com/products/json#sec-2.def
type Definition struct {
	VerbDivider string `json:"vd,omitempty"`
	WithSenseSequence
}
