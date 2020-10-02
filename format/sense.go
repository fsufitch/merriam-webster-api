package mwfmt

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/fsufitch/merriam-webster-api/types"
)

// SenseFormatter formats a sense
type SenseFormatter types.Sense

// ANSI fills the Formatter interface to create ANSI CLI output
func (f SenseFormatter) ANSI(opts FormatterOptions) (string, error) {
	builder := new(strings.Builder)

	prefix := opts.Prefix
	if len(f.SenseNumber) > 0 {
		if unicode.IsLetter([]rune(f.SenseNumber)[0]) {
			prefix += "  "
		}
		if unicode.IsPunct([]rune(f.SenseNumber)[0]) {
			prefix += "    "
		}
	}

	text, err := Format(f.DefiningText).ANSI(opts)
	if err != nil {
		return "", err
	}

	fmt.Fprintf(builder, "%s%s%s\n", prefix, opts.A.Bold(f.SenseNumber), text)

	return builder.String(), nil
}

// Plain fills the Formatter interface to create plaintext output
func (f SenseFormatter) Plain(opts FormatterOptions) (string, error) {
	return notImplementedString("SenseFormatter", "Plain")
}

// JSON fills the Formatter interface to create JSON output
func (f SenseFormatter) JSON(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("SenseFormatter", "JSON")
}

// HTML fills the Formatter interface to create HTML output
func (f SenseFormatter) HTML(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("SenseFormatter", "HTML")
}
