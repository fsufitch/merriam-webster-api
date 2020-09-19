package types

// WithVerbalIllustrations is a compositing type for parsing the `vis` property
type WithVerbalIllustrations struct {
	VerbalIllustrations []VerbalIllustration `json:"vis"`
}

// VerbalIllustration https://dictionaryapi.com/products/json#sec-2.vis
type VerbalIllustration struct {
	Text string `json:"t"`
	WithQuoteAttribution
}
