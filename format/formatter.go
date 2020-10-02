package mwfmt

import (
	"fmt"
	"reflect"

	"github.com/fsufitch/merriam-webster-api/types"
)

// Formatter is an interface for abstracting fromatting for different elements
type Formatter interface {
	ANSI(FormatterOptions) (string, error)
	Plain(FormatterOptions) (string, error)
	JSON(FormatterOptions) ([]byte, error)
	HTML(FormatterOptions) ([]byte, error)
}

// Format picks the appropriate formatter for the given item. Item must be of a type defined in mwtypes.
func Format(item interface{}) Formatter {
	switch it := item.(type) {
	case []types.CollegiateResult:
		return CollegiateResultsFormatter(it)
	case types.CollegiateResult:
		return CollegiateResultFormatter(it)
	case []types.Pronounciation:
		return PronounciationsFormatter(it)
	case []types.Definition:
		return DefinitionsFormatter(it)
	case types.SenseSequence:
		return SenseSequenceFormatter(it)
	case types.Sense:
		return SenseFormatter(it)
	case types.DefiningText:
		return DefiningTextFormatter(it)
	default:
		fmt.Println(it, reflect.TypeOf(it))
		return notImplementedFormatter{reflect.TypeOf(it).Name()}
	}
}

type notImplementedFormatter struct{ what string }

func (f notImplementedFormatter) ANSI(opts FormatterOptions) (string, error) {
	return "", fmt.Errorf("formatter not found for %s", f.what)
}

func (f notImplementedFormatter) Plain(opts FormatterOptions) (string, error) {
	return "", fmt.Errorf("formatter not found for %s", f.what)
}

func (f notImplementedFormatter) JSON(opts FormatterOptions) ([]byte, error) {
	return nil, fmt.Errorf("formatter not found for %s", f.what)
}

func (f notImplementedFormatter) HTML(opts FormatterOptions) ([]byte, error) {
	return nil, fmt.Errorf("formatter not found for %s", f.what)
}

func notImplementedString(what, format string) (string, error) {
	return "", fmt.Errorf("format `%s` not implemented for %s", format, what)
}

func notImplementedByteArray(what, format string) ([]byte, error) {
	return nil, fmt.Errorf("format `%s` not implemented for %s", format, what)
}

// Temporary copypaste area:

// // ANSI fills the Formatter interface to create ANSI CLI output
// func (f XXXXX) ANSI(opts FormatterOptions) (string, error) {
// 	return notImplementedString("XXXXX", "ANSI")
// }

// // Plain fills the Formatter interface to create plaintext output
// func (f XXXXX) Plain(opts FormatterOptions) (string, error) {
// 	return notImplementedString("XXXXX", "Plain")
// }

// // JSON fills the Formatter interface to create JSON output
// func (f XXXXX) JSON(opts FormatterOptions) ([]byte, error) {
// 	return notImplementedByteArray("XXXXX", "JSON")
// }

// // HTML fills the Formatter interface to create HTML output
// func (f XXXXX) HTML(opts FormatterOptions) ([]byte, error) {
// 	return notImplementedByteArray("XXXXX", "HTML")
// }
