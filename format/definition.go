package mwfmt

import (
	"fmt"
	"strings"

	"github.com/fsufitch/merriam-webster-api/types"
)

// DefinitionsFormatter formats []Definition
type DefinitionsFormatter []types.Definition

// ANSI fills the Formatter interface to create ANSI CLI output
func (f DefinitionsFormatter) ANSI(opts FormatterOptions) (string, error) {
	builder := new(strings.Builder)
	for _, def := range f {
		if def.VerbDivider != "" {
			fmt.Fprintf(builder, "%s%s\n", opts.Prefix, opts.A.Italic(def.VerbDivider))
		}

		text, err := Format(*def.SenseSequence).ANSI(opts.WithPrefix(opts.Prefix + "  "))
		if err != nil {
			return "", err
		}
		fmt.Fprintf(builder, "%s", text)
	}
	return builder.String(), nil
}

// Plain fills the Formatter interface to create plaintext output
func (f DefinitionsFormatter) Plain(opts FormatterOptions) (string, error) {
	return notImplementedString("DefinitionsFormatter", "Plain")
}

// JSON fills the Formatter interface to create JSON output
func (f DefinitionsFormatter) JSON(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("DefinitionsFormatter", "JSON")
}

// HTML fills the Formatter interface to create HTML output
func (f DefinitionsFormatter) HTML(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("DefinitionsFormatter", "HTML")
}
