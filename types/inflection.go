package types

// WithInflections is a compositing type for parsing the `ins` property
type WithInflections struct {
	Inflections []Inflection `json:"ins,omitempty"`
}

// Inflection https://dictionaryapi.com/products/json#sec-2.ins
type Inflection struct {
	Spelled string `json:"if,omitempty"`
	Cutback string `json:"ifc,omitempty"`
	Label   string `json:"il,omitempty"`
	WithPronounciations
	WithSenseSpecificInflectionPluralLabel
}
