package types

import "github.com/pkg/errors"

// WithUndefinedRunonText is a compositing type for parsing the `utxt` property
type WithUndefinedRunonText struct {
	UndefinedRunonText *UndefinedRunonText `json:"utxt,omitempty"`
}

// UndefinedRunonText https://dictionaryapi.com/products/json#sec-2.uros
type UndefinedRunonText SequenceMapping

// UndefinedRunonTextItemType is an enum type for the types of items in UndefinedRunonText
type UndefinedRunonTextItemType int

// Values for UndefinedRunonTextItemType
const (
	UndefinedRunonTextItemTypeUnknown = iota
	UndefinedRunonTextItemTypeVerbalIllustrations
	UndefinedRunonTextItemTypeUsageNotes
)

// UndefinedRunonTextItemTypeFromString returns a UndefinedRunonTextItemType from its string ID
func UndefinedRunonTextItemTypeFromString(id string) UndefinedRunonTextItemType {
	switch id {
	case "vis":
		return UndefinedRunonTextItemTypeVerbalIllustrations
	case "uns":
		return UndefinedRunonTextItemTypeUsageNotes
	default:
		return UndefinedRunonTextItemTypeUnknown
	}
}

func (t UndefinedRunonTextItemType) String() string {
	return []string{"", "vis", "uns"}[t]
}

// Contents returns a copied slice of the contents in the UndefinedRunonText
func (uro UndefinedRunonText) Contents() ([]UndefinedRunonTextItem, error) {
	items := []UndefinedRunonTextItem{}
	for _, el := range uro {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := UndefinedRunonTextItemTypeFromString(key)
		switch typ {
		case UndefinedRunonTextItemTypeVerbalIllustrations:
			var out []VerbalIllustration
			err = el.UnmarshalValue(&out)
			items = append(items, UndefinedRunonTextItem{Type: typ, WithVerbalIllustrations: WithVerbalIllustrations{out}})
		case UndefinedRunonTextItemTypeUsageNotes:
			var out []UsageNote
			err = el.UnmarshalValue(&out)
			items = append(items, UndefinedRunonTextItem{Type: typ, WithUsageNotes: WithUsageNotes{out}})
		default:
			err = errors.New("unknown item type in verbal illustration")
		}
	}
	return items, nil
}

// UndefinedRunonTextItem is an item of the SI container
type UndefinedRunonTextItem struct {
	Type UndefinedRunonTextItemType
	WithVerbalIllustrations
	WithUsageNotes
}
