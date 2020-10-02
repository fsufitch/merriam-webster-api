package mwfmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXmlify(t *testing.T) {
	output := xmlify(`testing{b}xmlification{/b}of{bc}{sup}{ldquo}this {weird}{dxt|first|second||fourth|||seventh|}stuff{rdquo}{/sup}`)
	assert.Equal(t, `testing<Bold>xmlification</Bold>of<BoldColon /><Superscript><LeftQuote />this {weird}<XRToken nargs="8" arg0="first" arg1="second" arg2="" arg3="fourth" arg4="" arg5="" arg6="seventh" arg7="" />stuff<RightQuote /></Superscript>`, output)
}
