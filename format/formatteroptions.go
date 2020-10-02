package mwfmt

import "github.com/logrusorgru/aurora/v3"

// FormatterOptions stores and transfers state between formatters
type FormatterOptions struct {
	A          aurora.Aurora
	Prefix     string
	Homographs int
}

// WithAurora copies the FormatterOptions but with the appropriate value
func (fo FormatterOptions) WithAurora(a aurora.Aurora) FormatterOptions {
	fo.A = a
	return fo
}

// WithPrefix copies the FormatterOptions but with the appropriate value
func (fo FormatterOptions) WithPrefix(p string) FormatterOptions {
	fo.Prefix = p
	return fo
}

// WithHomographs the FormatterOptions but with the appropriate value
func (fo FormatterOptions) WithHomographs(h int) FormatterOptions {
	fo.Homographs = h
	return fo
}
