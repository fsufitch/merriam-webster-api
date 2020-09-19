package types

import "github.com/pkg/errors"

// WithUsageNotes is a compositing type for parsing the `uns` property
type withUsageNotes struct {
	UsageNotes []UsageNote
}

// UsageNote // https://dictionaryapi.com/products/json#sec-2.uns
type UsageNote SequenceMapping

// UsageNoteItemType is an enum type for the types of items in the Usage Note
type UsageNoteItemType int

// Values for UsageNoteItemType
const (
	UsageNoteItemTypeUnknown UsageNoteItemType = iota
	UsageNoteItemTypeText
	UsageNoteItemTypeRunIn
	UsageNoteItemTypeVerbalIllustration
)

// UsageNoteItemTypeFromString returns a UsageNoteItemType from its string ID
func UsageNoteItemTypeFromString(id string) UsageNoteItemType {
	switch id {
	case "text":
		return UsageNoteItemTypeText
	case "ri":
		return UsageNoteItemTypeRunIn
	case "vis":
		return UsageNoteItemTypeVerbalIllustration
	default:
		return UsageNoteItemTypeUnknown
	}
}

func (t UsageNoteItemType) String() string {
	return []string{"", "text", "ri", "vis"}[t]
}

// Contents returns a copied slice of the contents in the UsageNote
func (un UsageNote) Contents() ([]UsageNoteItem, error) {
	items := []UsageNoteItem{}
	for _, el := range un {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := UsageNoteItemTypeFromString(key)
		switch typ {
		case SupplementalInfoItemTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			items = append(items, UsageNoteItem{Type: typ, Text: &out})
		case SupplementalInfoItemTypeRunIn:
			var out RunIn
			err = el.UnmarshalValue(&out)
			items = append(items, UsageNoteItem{Type: typ, RunIn: &out})
		case SupplementalInfoItemTypeVerbalIllustration:
			var out VerbalIllustration
			err = el.UnmarshalValue(&out)
			items = append(items, UsageNoteItem{Type: typ, VerbalIllustration: &out})
		default:
			err = errors.New("unknown item type in supplemental info")
		}
	}
	return items, nil
}

// UsageNoteItem is an item of the UsageNote container
type UsageNoteItem struct {
	Type               UsageNoteItemType
	Text               *string
	RunIn              *RunIn
	VerbalIllustration *VerbalIllustration
}
