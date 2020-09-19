package types

// WithCognateCrossReferences is a compositing type for parsing the `cxs` property
type WithCognateCrossReferences struct {
	CognateCrossReferences []CognateCrossReference `json:"cxs,omitempty"`
}

// CognateCrossReference https://dictionaryapi.com/products/json#sec-2.cxs
type CognateCrossReference struct {
	Label   string                        `json:"cxl,omitempty"`
	Targets []CognateCrossReferenceTarget `json:"cxtis,omitempty"`
}

// CognateCrossReferenceTarget https://dictionaryapi.com/products/json#sec-2.cxs
type CognateCrossReferenceTarget struct {
	Label         string `json:"cxl,omitempty"`
	TargetID      string `json:"cxr,omitempty"`
	HyperlinkText string `json:"cxt,omitempty"`
	SenseNumber   string `json:"cxn,omitempty"`
}
