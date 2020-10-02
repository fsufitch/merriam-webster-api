package mwfmt

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fsufitch/merriam-webster-api/types"
)

// CollegiateResultsFormatter formats []CollegiateResult
type CollegiateResultsFormatter []types.CollegiateResult

// ANSI fills the Formatter interface to create ANSI CLI output
func (f CollegiateResultsFormatter) ANSI(opts FormatterOptions) (string, error) {
	if len(f) < 1 {
		return "", errors.New("no results to format")
	}
	hw := strings.SplitN(f[0].Metadata.ID, ":", 2)[0]

	homographs := []types.CollegiateResult{f[0]}
	otherDefs := []types.CollegiateResult{}
	for _, result := range f[1:] {
		if strings.HasPrefix(result.Metadata.ID, hw+":") {
			homographs = append(homographs, result)
		} else {
			otherDefs = append(otherDefs, result)
		}
	}

	builder := new(strings.Builder)

	for _, result := range homographs {
		formatter := Format(result)

		var text string
		var err error
		if text, err = formatter.ANSI(opts.WithHomographs(len(homographs))); err != nil {
			return "", err
		}
		fmt.Fprintf(builder, "%s\n\n", text)
	}

	if len(otherDefs) > 0 {
		fmt.Fprint(builder, opts.A.Bold("Related search terms: "))

		terms := []string{}
		termsFound := map[string]struct{}{}
		for _, result := range f {
			term := strings.SplitN(result.Metadata.ID, ":", 2)[0]
			if _, ok := termsFound[term]; !ok {
				termsFound[term] = struct{}{}
				terms = append(terms, opts.A.Underline(opts.A.Blue(term)).String())
			}
		}

		fmt.Fprintf(builder, "%s\n\n", strings.Join(terms, ", "))
	}

	return builder.String(), nil
}

// Plain fills the Formatter interface to create plaintext output
func (f CollegiateResultsFormatter) Plain(opts FormatterOptions) (string, error) {
	return notImplementedString("[]CollegiateResult", "Plain")
}

// JSON fills the Formatter interface to create JSON output
func (f CollegiateResultsFormatter) JSON(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("[]CollegiateResult", "JSON")
}

// HTML fills the Formatter interface to create HTML output
func (f CollegiateResultsFormatter) HTML(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("[]CollegiateResult", "HTML")
}

// ==============================

// CollegiateResultFormatter formats CollegiateResult
type CollegiateResultFormatter types.CollegiateResult

// ANSI fills the Formatter interface to create ANSI CLI output
func (f CollegiateResultFormatter) ANSI(opts FormatterOptions) (string, error) {
	builder := new(strings.Builder)
	prefix := "  "

	hw := strings.ReplaceAll(f.HeadwordInfo.Headword, "*", "")
	hwsep := strings.ReplaceAll(f.HeadwordInfo.Headword, "*", "·")

	fmt.Fprintf(builder, "%s %s", opts.A.Underline(opts.A.Bold(hw)), opts.A.Italic(f.Function))
	if opts.Homographs > 1 {
		fmt.Fprintf(builder, opts.A.Faint(" (%d/%d)").String(), f.Homograph, opts.Homographs)
	}
	fmt.Fprintln(builder)

	if len(f.HeadwordInfo.Pronounciations) > 0 {
		fmt.Fprintf(builder, "%s%s", prefix, hwsep)
		prs, err := Format(f.HeadwordInfo.Pronounciations).ANSI(opts)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(builder, " | %s\n", prs)
	}

	defs, err := Format(f.Definitions).ANSI(opts.WithPrefix("  "))
	if err != nil {
		return "", err
	}
	fmt.Fprintf(builder, "%s", defs)

	return builder.String(), nil
}

// Plain fills the Formatter interface to create plaintext output
func (f CollegiateResultFormatter) Plain(opts FormatterOptions) (string, error) {
	return notImplementedString("CollegiateResult", "Plain")
}

// JSON fills the Formatter interface to create JSON output
func (f CollegiateResultFormatter) JSON(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("CollegiateResult", "JSON")
}

// HTML fills the Formatter interface to create HTML output
func (f CollegiateResultFormatter) HTML(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("CollegiateResult", "HTML")
}
