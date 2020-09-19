package types

import "github.com/pkg/errors"

// WithUsages is a compositing type for parsing the `usages` property
type WithUsages struct {
	// https://dictionaryapi.com/products/json#sec-2.usages
	UsageParagraphs []UsageParagraphs `json:"usages,omitempty"`
}

// UsageParagraphs https://dictionaryapi.com/products/json#sec-2.usages
type UsageParagraphs struct {
	Label string              `json:"pl,omitempty"`
	Text  *UsageParagraphText `json:"pt,omitempty"`
}

// UsageParagraphText https://dictionaryapi.com/products/json#sec-2.usages
type UsageParagraphText SequenceMapping

// UsageParagraphTextItemType is an enum type for the types of items in the paragraph
type UsageParagraphTextItemType int

// Values for UsageParagraphTextItemType
const (
	UsageParagraphTextItemTypeUnknown UsageParagraphTextItemType = iota
	UsageParagraphTextItemTypeText
	UsageParagraphTextItemTypeVerbalIllustration
	UsageParagraphTextItemTypeSeeAlso
)

// UsageParagraphTextItemTypeFromString returns a UsageParagraphTextItemType from its string ID
func UsageParagraphTextItemTypeFromString(id string) UsageParagraphTextItemType {
	switch id {
	case "text":
		return UsageParagraphTextItemTypeText
	case "vis":
		return UsageParagraphTextItemTypeVerbalIllustration
	case "uarefs":
		return UsageParagraphTextItemTypeSeeAlso
	default:
		return UsageParagraphTextItemTypeUnknown
	}
}

func (t UsageParagraphTextItemType) String() string {
	return []string{"", "text", "vis", "uarefs"}[t]
}

// Contents returns a copied slice of the contents in the UsageParagraphText
func (upt UsageParagraphText) Contents() ([]UsageParagraphTextItem, error) {
	items := []UsageParagraphTextItem{}
	for _, el := range upt {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := UsageParagraphTextItemTypeFromString(key)
		switch typ {
		case UsageParagraphTextItemTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			items = append(items, UsageParagraphTextItem{Type: typ, Text: &out})
		case UsageParagraphTextItemTypeVerbalIllustration:
			var out VerbalIllustration
			err = el.UnmarshalValue(&out)
			items = append(items, UsageParagraphTextItem{Type: typ, VerbalIllustration: &out})
		case UsageParagraphTextItemTypeSeeAlso:
			var out []UsageSeeAlso
			err = el.UnmarshalValue(&out)
			items = append(items, UsageParagraphTextItem{Type: typ, SeeAlso: out})
		default:
			err = errors.New("unknown element type in run-in")
		}
	}
	return items, nil
}

// UsageParagraphTextItem is an item of the UsageParagraphText container
type UsageParagraphTextItem struct {
	Type               UsageParagraphTextItemType
	Text               *string
	VerbalIllustration *VerbalIllustration
	SeeAlso            []UsageSeeAlso
}

// UsageSeeAlso is "uaref" here https://dictionaryapi.com/products/json#sec-2.usages
type UsageSeeAlso struct {
	Reference string `json:"uaref,omitempty"`
}
