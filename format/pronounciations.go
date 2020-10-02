package mwfmt

import (
	"fmt"
	"strings"

	"github.com/fsufitch/merriam-webster-api/types"
)

// PronounciationsFormatter formats []Pronounciation
type PronounciationsFormatter []types.Pronounciation

// ANSI fills the Formatter interface to create ANSI CLI output
func (f PronounciationsFormatter) ANSI(opts FormatterOptions) (string, error) {
	pronounciations := []string{}
	punctuation := ", "
	for _, p := range f {
		if p.Punctuation != "" {
			punctuation = p.Punctuation
		}
		builder := new(strings.Builder)
		if p.LabelBefore != "" {
			fmt.Fprintf(builder, "%s ", opts.A.Italic(p.LabelBefore))
		}
		fmt.Fprintf(builder, "%s", p.MerriamWebsterFormat)
		if p.LabelAfter != "" {
			fmt.Fprintf(builder, " %s", opts.A.Italic(p.LabelAfter))
		}
		pronounciations = append(pronounciations, builder.String())
	}

	return fmt.Sprintf("\\ %s \\", strings.Join(pronounciations, punctuation)), nil
}

// Plain fills the Formatter interface to create plaintext output
func (f PronounciationsFormatter) Plain(opts FormatterOptions) (string, error) {
	return notImplementedString("PronounciationsFormatter", "Plain")
}

// JSON fills the Formatter interface to create JSON output
func (f PronounciationsFormatter) JSON(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("PronounciationsFormatter", "JSON")
}

// HTML fills the Formatter interface to create HTML output
func (f PronounciationsFormatter) HTML(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("PronounciationsFormatter", "HTML")
}
