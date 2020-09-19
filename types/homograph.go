package types

// WithHomograph is a compositing type for parsing the `hom` property
type WithHomograph struct {
	// https://dictionaryapi.com/products/json#sec-2.hom
	Homograph int `json:"hom"`
}
