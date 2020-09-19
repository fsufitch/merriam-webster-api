package types

// WithMetadata is a compositing type for parsing the `metadata` property
type WithMetadata struct {
	Metadata *Metadata `json:"meta,omitempty"`
}

// Metadata https://dictionaryapi.com/products/json#sec-2.meta
type Metadata struct {
	ID        string   `json:"id,omitempty"`
	UUID      string   `json:"uuid,omitempty"`
	SortKey   string   `json:"sort,omitempty"`
	Source    string   `json:"src,omitempty"`
	Section   string   `json:"section,omitempty"`
	Stems     []string `json:"stems,omitempty"`
	Offensive bool     `json:"offensive"`
}
