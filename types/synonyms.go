package types

import "github.com/pkg/errors"

// WithSynonyms is a compositing type for parsing the `syns` property
type WithSynonyms struct {
	// https://dictionaryapi.com/products/json#sec-2.syns
	SynonymParagraphs []SynonymParagraph `json:"syns"`
}

// SynonymParagraph https://dictionaryapi.com/products/json#sec-2.syns
type SynonymParagraph struct {
	Label       string               `json:"pl"`
	Text        SynonymParagraphText `json:"pt"`
	SeeAlsoRefs []string             `json:"sarefs"`
}

// SynonymParagraphText https://dictionaryapi.com/products/json#sec-2.syns
type SynonymParagraphText SequenceMapping

// SynonymParagraphTextItemType is an enum type for the types of items in the paragraph
type SynonymParagraphTextItemType int

// Values for SynonymParagraphTextItemType
const (
	SynonymParagraphTextItemTypeUnknown SynonymParagraphTextItemType = iota
	SynonymParagraphTextItemTypeText
	SynonymParagraphTextItemTypeVerbalIllustration
)

// SynonymParagraphTextItemTypeFromString returns a SynonymParagraphTextItemType from its string ID
func SynonymParagraphTextItemTypeFromString(id string) SynonymParagraphTextItemType {
	switch id {
	case "text":
		return SynonymParagraphTextItemTypeText
	case "vis":
		return SynonymParagraphTextItemTypeVerbalIllustration
	default:
		return SynonymParagraphTextItemTypeUnknown
	}
}

func (t SynonymParagraphTextItemType) String() string {
	return []string{"", "text", "vis"}[t]
}

// Contents returns a copied slice of the contents in the SynonymParagraphText
func (upt SynonymParagraphText) Contents() ([]SynonymParagraphTextItem, error) {
	items := []SynonymParagraphTextItem{}
	for _, el := range upt {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := SynonymParagraphTextItemTypeFromString(key)
		switch typ {
		case SynonymParagraphTextItemTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			items = append(items, SynonymParagraphTextItem{Type: typ, Text: &out})
		case SynonymParagraphTextItemTypeVerbalIllustration:
			var out VerbalIllustration
			err = el.UnmarshalValue(&out)
			items = append(items, SynonymParagraphTextItem{Type: typ, VerbalIllustration: &out})
		default:
			err = errors.New("unknown item type in run-in")
		}
	}
	return items, nil
}

// SynonymParagraphTextItem is an item of the SynonymParagraphText container
type SynonymParagraphTextItem struct {
	Type               SynonymParagraphTextItemType
	Text               *string
	VerbalIllustration *VerbalIllustration
}
