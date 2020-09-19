package types

// WithDirectionalCrossReferences is a compositing type for parsing the `dxlns` property
type WithDirectionalCrossReferences struct {
	DirectionalCrossReferences []string `json:"dxnls"`
}
