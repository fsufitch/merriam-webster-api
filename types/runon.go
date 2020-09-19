package types

// WithDefinedRunOns is a compositing type for parsing the `dros` property
type WithDefinedRunOns struct {
	// https://dictionaryapi.com/products/json#sec-2.dros
	DefinedRunOns []DefinedRunOn `json:"dros,omitempty"`
}

// DefinedRunOn https://dictionaryapi.com/products/json#sec-2.dros
type DefinedRunOn struct {
	Phrase string `json:"drp,omitempty"`
	WithDefinitions
	WithEtymology
	WithGeneralLabels
	WithPronounciations
	WithParenthesizedSubjectStatusLabel
	WithSubjectStatusLabels
	WithVariants
}

// WithUndefinedRunOns is a compositing type for parsing the `uro` property
type WithUndefinedRunOns struct {
	// https://dictionaryapi.com/products/json#sec-2.uros
	UndefinedRunOns []UndefinedRunOn `json:"uro,omitempty"`
}

// UndefinedRunOn https://dictionaryapi.com/products/json#sec-2.uros
type UndefinedRunOn struct {
	Word string `json:"ure,omitempty"`
	WithFunctionalLabel
	WithUndefinedRunonText
}
