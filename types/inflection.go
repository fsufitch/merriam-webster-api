package types

// WithInflections is a compositing type for parsing the `ins` property
type WithInflections struct {
	Inflections []Inflection `json:"ins"`
}

// Inflection https://dictionaryapi.com/products/json#sec-2.ins
type Inflection struct {
	Spelled string `json:"if"`
	Cutback string `json:"ifc"`
	Label   string `json:"il"`
	WithPronounciations
	WithSenseSpecificInflectionPluralLabel
}
