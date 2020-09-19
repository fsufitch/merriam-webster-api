package types

// WithQuoteAttribution is a compositing type for parsing the `aq` property
type WithQuoteAttribution struct {
	// https://dictionaryapi.com/products/json#sec-2.aq
	QuoteAttribution *QuoteAttribution `json:"aq,omitempty"`
}

// QuoteAttribution https://dictionaryapi.com/products/json#sec-2.aq
type QuoteAttribution struct {
	Author    string `json:"auth,omitempty"`
	Source    string `json:"source,omitempty"`
	Date      string `json:"aqdate,omitempty"`
	SubSource *struct {
		Source string `json:"source,omitempty"`
		Date   string `json:"aqdate,omitempty"`
	} `json:"subsource,omitempty"`
}
