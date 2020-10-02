package mwfmt

import (
	"fmt"
	"strings"

	"github.com/fsufitch/merriam-webster-api/types"
)

// SenseSequenceFormatter formats a SenseSequence
type SenseSequenceFormatter types.SenseSequence

// ANSI fills the Formatter interface to create ANSI CLI output
func (f SenseSequenceFormatter) ANSI(opts FormatterOptions) (string, error) {
	items, err := types.SenseSequence(f).Contents()

	if err != nil {
		return "", err

	}
	builder := new(strings.Builder)
	nextOpts := opts.WithPrefix(opts.Prefix)
	for _, item := range items {
		switch item.Type {
		case types.SenseSequenceItemTypeSense:
			text, err := Format(*item.Sense).ANSI(nextOpts)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(builder, "%s", text)
		// case types.SenseSequenceItemTypeAbbreviatedSense:
		// case types.SenseSequenceItemTypeBindingSubstitute:
		case types.SenseSequenceItemTypeSubSequence:
			text, err := Format(*item.SubSequence).ANSI(nextOpts)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(builder, "%s", text)

		case types.SenseSequenceItemTypeParenthesizedSequence:
			text, err := Format(*item.ParenthesizedSequence).ANSI(nextOpts)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(builder, "%s", text)
		default:
			fmt.Fprintf(builder, opts.A.Yellow("%s{Unsupported sseq item type: %s}\n").String(), opts.Prefix, item.Type)
		}
	}

	return builder.String(), nil
}

// Plain fills the Formatter interface to create plaintext output
func (f SenseSequenceFormatter) Plain(opts FormatterOptions) (string, error) {
	return notImplementedString("SenseSequenceFormatter", "Plain")
}

// JSON fills the Formatter interface to create JSON output
func (f SenseSequenceFormatter) JSON(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("SenseSequenceFormatter", "JSON")
}

// HTML fills the Formatter interface to create HTML output
func (f SenseSequenceFormatter) HTML(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("SenseSequenceFormatter", "HTML")
}
