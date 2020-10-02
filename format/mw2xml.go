package mwfmt

import (
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func xmlify(s string) string {
	for _, tag := range tags {
		s = tag.XMLify(s)
	}

	return s
}

type mw2xmlTag struct {
	MW          string
	XML         string
	SelfClosing bool
	Regexps     *regexps
}
type regexps struct {
	Open  *regexp.Regexp
	Close *regexp.Regexp
}

var tags = []mw2xmlTag{
	// https://dictionaryapi.com/products/json#sec-2.fmttokens
	{"b", "Bold", false, nil},
	{"bc", "BoldColon", true, nil},
	{"inf", "Subscript", false, nil},
	{"it", "Italics", false, nil},
	{"ldquo", "LeftQuote", true, nil},
	{"rdquo", "RightQuote", true, nil},
	{"sc", "SmallCaps", false, nil},
	{"sup", "Superscript", false, nil},

	// https://dictionaryapi.com/products/json#sec-2.wordtokens
	{"gloss", "Gloss", false, nil},
	{"parahw", "ParagraphHeadword", false, nil},
	{"phrase", "Phrase", false, nil},
	{"qword", "QuoteHeadword", false, nil},
	{"wi", "Headword", false, nil},

	// https://dictionaryapi.com/products/json#sec-2.xrefregtokens
	{"dx", "XR", false, nil},
	{"dx_def", "XRParanthetical", false, nil},
	{"dx_ety", "XRDirectional", false, nil},
	{"ma", "XRMoreAt", false, nil},

	// https://dictionaryapi.com/products/json#sec-2.xreftokens
	{"a_link", "AutoLinkToken", true, nil},
	{"d_link", "DirectLinkToken", true, nil},
	{"i_link", "ItalicLinkToken", true, nil},
	{"et_link", "EtymologyLinkToken", true, nil},
	{"mat", "MoreAtToken", true, nil},
	{"sx", "XRSynonymousToken", true, nil},
	{"dxt", "XRToken", true, nil},

	// https://dictionaryapi.com/products/json#sec-2.dstoken
	{"ds", "DateSenseToken", true, nil},
}

func (t *mw2xmlTag) MustRegexps() {
	if t.Regexps != nil {
		return
	}
	t.Regexps = new(regexps)
	re1 := regexp.MustCompile(fmt.Sprintf(`\{%s(\|[\w\s\|:-]+)?\}`, t.MW))
	t.Regexps.Open = re1
	if !t.SelfClosing {
		t.Regexps.Close = regexp.MustCompile(fmt.Sprintf(`\{/%s\}`, t.MW))
	}
}

func (t *mw2xmlTag) XMLify(s string) string {
	t.MustRegexps()

	s = t.Regexps.Open.ReplaceAllStringFunc(s, func(match string) string {
		return t.BuildXMLTag(match, false)
	})

	if !t.SelfClosing {
		s = t.Regexps.Close.ReplaceAllStringFunc(s, func(match string) string {
			return t.BuildXMLTag(match, true)
		})
	}
	return s
}

func (t *mw2xmlTag) BuildXMLTag(match string, closing bool) string {
	if closing {
		return fmt.Sprintf("</%s>", t.XML)
	}

	args := strings.Split(t.Regexps.Open.FindStringSubmatch(match)[1], "|")[1:]

	builder := new(strings.Builder)
	fmt.Fprintf(builder, "<%s", t.XML)
	if len(args) > 0 {
		fmt.Fprintf(builder, ` nargs="%d"`, len(args))
	}
	for i, arg := range args {
		fmt.Fprintf(builder, ` arg%d="%s"`, i, xmlescape(arg))
	}

	switch t.SelfClosing {
	case true:
		fmt.Fprintf(builder, ` />`)
	case false:
		fmt.Fprintf(builder, `>`)
	}

	return builder.String()
}

func xmlescape(s string) string {
	buf := new(strings.Builder)
	if err := xml.EscapeText(buf, []byte(s)); err != nil {
		panic(err)
	}
	return buf.String()
}

func xmlunescape(s string) string {
	return html.UnescapeString(s)
}

// MWXMLNode is a node of the XML tree of a parsed MW text
type MWXMLNode struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",attr"`
	Children []MWXMLChild
}

// MWXMLChild contains either plain text or a child node in a MW text XML tree
type MWXMLChild struct {
	Text string
	Node *MWXMLNode
}

// UnmarshalXML is used to integrate with `encoding/xml`
func (p *MWXMLNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	p.XMLName = start.Copy().Name
	p.Attrs = start.Copy().Attr
	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch t := token.(type) {
		case xml.CharData:
			p.Children = append(p.Children, MWXMLChild{Text: xmlunescape(string(t))})
		case xml.StartElement:
			n := new(MWXMLNode)
			if err = d.DecodeElement(n, &t); err != nil {
				return err
			}
			p.Children = append(p.Children, MWXMLChild{Node: n})
		default:
			// Other token types (end tag, etc) are ignored
		}
	}
	return nil
}

// GetAttr retrieves an attr from a node. Returns (value, ok)
func (p MWXMLNode) GetAttr(attr string) (string, bool) {
	for _, a := range p.Attrs {
		if a.Name.Local == attr {
			return a.Value, true
		}
	}
	return "", false
}

// GetArgs returns the list of pipe-separated args in the original tag
func (p MWXMLNode) GetArgs() ([]string, error) {
	var (
		args  []string
		nargs int
		err   error
	)

	if nargsStr, ok := p.GetAttr("nargs"); !ok {
		return nil, errors.New("no `nargs` attr found on tag")
	} else if nargs, err = strconv.Atoi(nargsStr); err != nil {
		return nil, fmt.Errorf("value of nargs was not int (%v): %v", nargs, err)
	}

	for i := 0; i < nargs; i++ {
		var arg string
		var ok bool
		if arg, ok = p.GetAttr(fmt.Sprintf("arg%d", i)); !ok {
			return nil, fmt.Errorf("arg%d not found in node", i)
		}
		args = append(args, arg)
	}
	return args, nil
}
