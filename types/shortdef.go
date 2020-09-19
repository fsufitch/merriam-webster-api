package types

// WithShortDefinitions is a compositing type for parsing the `shortdef` property
type WithShortDefinitions struct {
	ShortDefinitions []string `json:"shortdef,omitempty"`
}
