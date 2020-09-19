package types

// WithArt is a compositing type for parsing the `art` property
type WithArt struct {
	// https://dictionaryapi.com/products/json#sec-2.art
	Art *Art `json:"art,omitempty"`
}

// Art https://dictionaryapi.com/products/json#sec-2.art
type Art struct {
	ID      string `json:"artid,omitempty"`
	Caption string `json:"capt,omitempty"`
}
