package types

// WithHeadwordInfo is a compositing type for parsing the `hwi` property
type WithHeadwordInfo struct {
	HeadwordInfo *HeadwordInfo `json:"hwi,omitempty"`
}

// HeadwordInfo https://dictionaryapi.com/products/json#sec-2.hwi
type HeadwordInfo struct {
	Headword string `json:"hw,omitempty"`
	WithPronounciations
}
