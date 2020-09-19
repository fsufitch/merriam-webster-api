package types

// WithPronounciations is a compositing type for parsing the `prs` property
type WithPronounciations struct {
	Pronounciations []Pronounciation `json:"prs"`
}

// Pronounciation https://dictionaryapi.com/products/json#sec-2.prs
type Pronounciation struct {
	MerriamWebsterFormat string `json:"mw"`
	LabelBefore          string `json:"l"`
	LabelAfter           string `json:"l2"`
	Punctuation          string `json:"pun"`
	Sound                struct {
		Filename string `json:"audio"`
	} `json:"sound"`
}

// Sound https://dictionaryapi.com/products/json#sec-2.prs
// To see how to use this sound file reference
type Sound struct {
	Filename string `json:"audio"`
}
