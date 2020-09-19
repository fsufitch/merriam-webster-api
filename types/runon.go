package types

// WithDefinedRunOns is a compositing type for parsing the `dros` property
type WithDefinedRunOns struct {
	// https://dictionaryapi.com/products/json#sec-2.dros
	DefinedRunOns []DefinedRunOn `json:"dros"`
}

// DefinedRunOn https://dictionaryapi.com/products/json#sec-2.dros
type DefinedRunOn struct {
	Phrase string `json:"drp"`
	WithDefinitions
	WithEtymology
	WithGeneralLabels
	WithPronounciations
	WithParenthesizedSubjectStatusLabel
	WithSubjectStatusLabels
	WithVariants
}

type withUndefinedRunOns struct {
	// https://dictionaryapi.com/products/json#sec-2.uros
}

// UndefinedRunOn https://dictionaryapi.com/products/json#sec-2.uros
type UndefinedRunOn struct {
	Word string `json:"ure"`
	WithFunctionalLabel
	Texts struct{} `json:"utxt"`
}
