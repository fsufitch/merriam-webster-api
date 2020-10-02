package mwfmt

import (
	"fmt"
	"strings"

	"github.com/fsufitch/merriam-webster-api/types"
)

// DefiningTextFormatter formats DefiningText
type DefiningTextFormatter types.DefiningText

// ANSI fills the Formatter interface to create ANSI CLI output
func (f DefiningTextFormatter) ANSI(opts FormatterOptions) (string, error) {
	items, err := types.DefiningText(f).Contents()
	if err != nil {
		return "", err
	}

	builder := new(strings.Builder)
	for _, item := range items {
		switch item.Type {
		case types.DefiningTextItemTypeText:
			text, err := MerriamWebsterTagTextFormatter(*item.Text).ANSI(opts)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(builder, "%s", text)

		default:
			fmt.Fprintf(builder, opts.A.Yellow("{Unknown dt item type: %s} ").String(), item.Type)
		}
	}
	return builder.String(), nil
}

// Plain fills the Formatter interface to create plaintext output
func (f DefiningTextFormatter) Plain(opts FormatterOptions) (string, error) {
	return notImplementedString("DefiningTextFormatter", "Plain")
}

// JSON fills the Formatter interface to create JSON output
func (f DefiningTextFormatter) JSON(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("DefiningTextFormatter", "JSON")
}

// HTML fills the Formatter interface to create HTML output
func (f DefiningTextFormatter) HTML(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("DefiningTextFormatter", "HTML")
}
