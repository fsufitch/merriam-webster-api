package mwfmt

import (
	"fmt"
	"testing"

	"github.com/logrusorgru/aurora/v3"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	text, err := MerriamWebsterTagTextFormatter(`
	hello {b}world{/b}! all work and no play {bc} makes jack a {it}dull {b}boy{/b}{/it}. 
	this is {inf}subtext{/inf} and this is {sup}supertext{/sup}
	this is an {a_link|autolink}
	this is a {d_link|direct link|} and one with a {d_link|destination|go here}
	this is an {i_link|italic link|} and one with a {i_link|destination|go here}
	this is an {et_link|etymology link|} and one with a {et_link|destination|go here}
	{ma}{mat|here|}{/ma}, or {ma}{mat|there|actually here}{/ma}
	this is a {sx|synonymous cross reference||}, or {sx|another|with a location|}, or {sx|yet another|elsewhere|with some text afterward}
	DX: a bird of any kind {dx}compare {dxt|waterfowl||}, {dxt|wildfowl||}{/dx}
	DX: a bird of any kind {dx}compare {dxt|waterfowl|jackdaw|}, {dxt|wildfowl|crow|}{/dx}
	DX: a bird of any kind {dx}compare {dxt|waterfowl|jackdaw|illustration}, {dxt|wildfowl|crow|illustration}{/dx}
	DX: a bird of any kind {dx}compare {dxt|table|jackdaw|table}, {dxt|stool|crow|table}{/dx}
	{ds|v|1|a|b}
	`).ANSI(FormatterOptions{A: aurora.NewAurora(true)})
	assert.NoError(t, err)
	fmt.Println(text)
}
