package types

// WithPronounciations is a compositing type for parsing the `prs` property
type WithPronounciations struct {
	Pronounciations []Pronounciation `json:"prs,omitempty"`
}

// Pronounciation https://dictionaryapi.com/products/json#sec-2.prs
type Pronounciation struct {
	MerriamWebsterFormat string `json:"mw,omitempty"`
	LabelBefore          string `json:"l,omitempty"`
	LabelAfter           string `json:"l2,omitempty"`
	Punctuation          string `json:"pun,omitempty"`
	Sound                *Sound `json:"sound,omitempty"`
}

// Sound https://dictionaryapi.com/products/json#sec-2.prs
// To see how to use this sound file reference
type Sound struct {
	Filename string `json:"audio,omitempty"`
	Ref      string `json:"ref,omitempty"`
	Stat     string `json:"stat,omitempty"`
}
