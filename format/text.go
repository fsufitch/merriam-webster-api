package mwfmt

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"

	aurora "github.com/logrusorgru/aurora/v3"
)

// MerriamWebsterTagTextFormatter formats text that contains MW tokens
// https://dictionaryapi.com/products/json#sec-2.tokens
type MerriamWebsterTagTextFormatter string

// ANSI fills the Formatter interface to create ANSI CLI output
func (f MerriamWebsterTagTextFormatter) ANSI(opts FormatterOptions) (string, error) {
	xmlified := fmt.Sprintf("<ParsedText>%s</ParsedText>", xmlify(string(f)))
	el := new(MWXMLNode)
	if err := xml.Unmarshal([]byte(xmlified), el); err != nil {
		return "", err
	}

	return renderNodeANSI(el, opts.A)
}

// Plain fills the Formatter interface to create plaintext output
func (f MerriamWebsterTagTextFormatter) Plain(opts FormatterOptions) (string, error) {
	return notImplementedString("MerriamWebsterTagTextFormatter", "Plain")
}

// JSON fills the Formatter interface to create JSON output
func (f MerriamWebsterTagTextFormatter) JSON(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("MerriamWebsterTagTextFormatter", "JSON")
}

// HTML fills the Formatter interface to create HTML output
func (f MerriamWebsterTagTextFormatter) HTML(opts FormatterOptions) ([]byte, error) {
	return notImplementedByteArray("MerriamWebsterTagTextFormatter", "HTML")
}

func renderNodeANSI(n *MWXMLNode, a aurora.Aurora) (string, error) {
	builder := strings.Builder{}
	for _, child := range n.Children {
		if child.Node != nil {
			if renderedChild, err := renderNodeANSI(child.Node, a); err == nil {
				builder.WriteString(renderedChild)
			} else {
				return "", err
			}
		} else {
			builder.WriteString(child.Text)
		}
	}

	var formatted fmt.Stringer

	switch n.XMLName.Local {
	case "ParsedText":
		formatted = &builder

	// https://dictionaryapi.com/products/json#sec-2.fmttokens
	case "Bold":
		formatted = a.Bold(builder.String())
	case "BoldColon":
		formatted = a.Bold(" : ")
	case "Subscript":
		// TODO: something cleverer?
		formatted = a.Faint(a.Sprintf("_[%s]", builder.String()))
	case "Italics":
		formatted = a.Italic(builder.String())
	case "LeftQuote":
		formatted = a.Reset("\u201c")
	case "RightQuote":
		formatted = a.Reset("\u201d")
	case "SmallCaps":
		formatted = a.Gray(12, strings.ToUpper(builder.String()))
	case "Superscript":
		// TODO: something cleverer?
		formatted = a.Faint(a.Sprintf("^[%s]", builder.String()))

	// https://dictionaryapi.com/products/json#sec-2.wordtokens
	case "Gloss":
		formatted = bytes.NewBufferString(fmt.Sprintf("[%s]", builder.String()))
	case "ParagraphHeadword":
		// "bold smallcaps"
		formatted = a.Bold(a.Gray(12, strings.ToUpper(builder.String())))
	case "Phrase":
		formatted = a.Bold(a.Italic(builder.String()))
	case "QuoteHeadword", "Headword":
		formatted = a.Italic(builder.String())

	// https://dictionaryapi.com/products/json#sec-2.xrefregtokens
	case "XR", "XRDirectional":
		formatted = bytes.NewBufferString(fmt.Sprintf("— %s", builder.String()))
	case "XRMoreAt":
		formatted = bytes.NewBufferString(fmt.Sprintf("— more at %s", builder.String()))
	case "XRParanthetical":
		formatted = bytes.NewBufferString(fmt.Sprintf("(%s)", builder.String()))

	// https://dictionaryapi.com/products/json#sec-2.xreftokens
	case "AutoLinkToken":
		args, err := n.GetArgs()
		if err != nil {
			return "", fmt.Errorf("error rendering AutoLinkToken: %v", err)
		}
		if len(args) != 1 {
			return "", fmt.Errorf("AutoLinkToken contains unexpected number of args %d (expected 1)", len(args))
		}
		formatted = a.Blue(a.Underline(args[0]))
	case "DirectLinkToken":
		args, err := n.GetArgs()
		if err != nil {
			return "", fmt.Errorf("error rendering DirectLinkToken: %v", err)
		}
		if len(args) != 2 {
			return "", fmt.Errorf("DirectLinkToken contains unexpected number of args %d (expected 2)", len(args))
		}
		builder := new(strings.Builder)
		builder.WriteString(args[0])
		if args[1] != "" {
			fmt.Fprintf(builder, "(🔗%s)", args[1])
		}
		formatted = a.Blue(a.Underline(builder.String()))
	case "ItalicLinkToken":
		args, err := n.GetArgs()
		if err != nil {
			return "", fmt.Errorf("error rendering ItalicLinkToken: %v", err)
		}
		if len(args) != 2 {
			return "", fmt.Errorf("ItalicLinkToken contains unexpected number of args %d (expected 2)", len(args))
		}
		builder := new(strings.Builder)
		builder.WriteString(args[0])
		if args[1] != "" {
			fmt.Fprintf(builder, " (🔗%s)", args[1])
		}
		formatted = a.Italic(a.Blue(a.Underline(builder.String())))
	case "EtymologyLinkToken", "MoreAtToken":
		args, err := n.GetArgs()
		if err != nil {
			return "", fmt.Errorf("error rendering EtymologyLinkToken: %v", err)
		}
		if len(args) != 2 {
			return "", fmt.Errorf("EtymologyLinkToken contains unexpected number of args %d (expected 2)", len(args))
		}
		builder := new(strings.Builder)
		builder.WriteString(strings.ToUpper(args[0]))
		if args[1] != "" {
			fmt.Fprintf(builder, " (🔗%s)", args[1])
		}
		formatted = a.Faint(a.Blue(a.Underline(builder.String())))
	case "XRSynonymousToken":
		args, err := n.GetArgs()
		if err != nil {
			return "", fmt.Errorf("error rendering XRSynonymousToken: %v", err)
		}
		if len(args) != 3 {
			return "", fmt.Errorf("XRSynonymousToken contains unexpected number of args %d (expected 3)", len(args))
		}
		builder1 := new(strings.Builder)
		builder1.WriteString(strings.ToUpper(args[0]))
		if args[1] != "" {
			fmt.Fprintf(builder1, " (🔗%s)", args[1])
		}
		builder2 := new(strings.Builder)
		fmt.Fprint(builder2, a.Faint(a.Blue(a.Underline(builder1.String()))).String())
		if args[2] != "" {
			fmt.Fprintf(builder2, " %s", args[2])
		}
		formatted = builder2
	case "XRToken":
		args, err := n.GetArgs()
		if err != nil {
			return "", fmt.Errorf("error rendering XRToken: %v", err)
		}
		if len(args) != 3 {
			return "", fmt.Errorf("XRToken contains unexpected number of args %d (expected 3)", len(args))
		}
		link := args[0]
		target := args[1]
		after := args[2]
		if args[1] == "" {
			target = args[2]
		}
		if strings.Contains(args[2], "table") {
			after = ""
		}

		linkBuilder := new(strings.Builder)
		fmt.Fprintf(linkBuilder, "%s", link)
		if target != "" {
			fmt.Fprintf(linkBuilder, " (🔗%s)", target)
		}
		if after != "" {
			after = fmt.Sprintf(" %s", after)
		}

		formatted = bytes.NewBufferString(fmt.Sprintf("%s%s", a.Faint(a.Blue(a.Underline(linkBuilder.String()))).String(), after))

		// https://dictionaryapi.com/products/json#sec-2.dstoken
	case "DateSenseToken":
		args, err := n.GetArgs()
		if err != nil {
			return "", fmt.Errorf("error rendering DateSenseToken: %v", err)
		}
		if len(args) != 4 {
			return "", fmt.Errorf("DateSenseToken contains unexpected number of args %d (expected 4)", len(args))
		}
		if args[0] == "" && args[1] == "" && args[2] == "" && args[3] == "" {
			formatted = new(bytes.Buffer)
			break
		}
		builder := new(strings.Builder)
		fmt.Fprint(builder, "in the meaning defined at ")
		if args[0] != "" {
			fmt.Fprintf(builder, "(%s)", a.Faint(args[0]))
		}
		if args[1] != "" {
			fmt.Fprintf(builder, "%s", a.Bold(args[1]))
		}
		if args[2] != "" {
			fmt.Fprintf(builder, "%s", args[2])
		}
		if args[3] != "" {
			fmt.Fprintf(builder, "(%s)", args[3])
		}

		formatted = builder

	default:
		return "", fmt.Errorf("unrecognized XML tag `%s` {attrs: %+v} {children: %+v}", n.XMLName.Local, n.Attrs, n.Children)
	}
	return formatted.String(), nil
}
