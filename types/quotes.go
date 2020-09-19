package types

// WithQuotes is a compositing type for parsing the `quotes` property
type WithQuotes struct {
	// https://dictionaryapi.com/products/json#sec-2.quotes
	Quotes []Quote `json:"quotes,omitempty"`
}

// Quote https://dictionaryapi.com/products/json#sec-2.quotes
type Quote struct {
	Text string `json:"t,omitempty"`
	WithQuoteAttribution
}
