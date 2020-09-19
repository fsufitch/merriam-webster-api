package types

// WithFirstKnownDate is a compositing type for parsing the `date` property
type WithFirstKnownDate struct {
	FirstKnownDate string `json:"date,omitempty"`
}
