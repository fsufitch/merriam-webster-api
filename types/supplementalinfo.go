package types

import (
	"fmt"
)

// SupplementalInfo https://dictionaryapi.com/products/json#sec-2.snote
type SupplementalInfo SequenceMapping

// SupplementalInfoItemType is an enum type for the types of items in SupplementalInfo
type SupplementalInfoItemType int

// Values for SupplementalInfoItemType
const (
	SupplementalInfoItemTypeUnknown = iota
	SupplementalInfoItemTypeText
	SupplementalInfoItemTypeRunIn
	SupplementalInfoItemTypeVerbalIllustration
)

// SupplementalInfoItemTypeFromString returns a SupplementalInfoItemType from its string ID
func SupplementalInfoItemTypeFromString(id string) SupplementalInfoItemType {
	switch id {
	case "t":
		return SupplementalInfoItemTypeText
	case "ri":
		return SupplementalInfoItemTypeRunIn
	case "vis":
		return SupplementalInfoItemTypeVerbalIllustration
	default:
		return SupplementalInfoItemTypeUnknown
	}
}

func (t SupplementalInfoItemType) String() string {
	return []string{"", "t", "ri", "vis"}[t]
}

// Contents returns a copied slice of the contents in the SupplementalInfo
func (si SupplementalInfo) Contents() ([]SupplementalInfoItem, error) {
	items := []SupplementalInfoItem{}
	for _, el := range si {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := SupplementalInfoItemTypeFromString(key)
		switch typ {
		case SupplementalInfoItemTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			items = append(items, SupplementalInfoItem{Type: typ, Text: &out})
		case SupplementalInfoItemTypeRunIn:
			var out RunIn
			err = el.UnmarshalValue(&out)
			items = append(items, SupplementalInfoItem{Type: typ, RunIn: &out})
		case SupplementalInfoItemTypeVerbalIllustration:
			var out VerbalIllustration
			err = el.UnmarshalValue(&out)
			items = append(items, SupplementalInfoItem{Type: typ, VerbalIllustration: &out})
		default:
			err = fmt.Errorf("unknown item type in supplemental info: %v (enum: %v)", key, typ)
		}
	}
	return items, nil
}

// SupplementalInfoItem is an item of the SI container
type SupplementalInfoItem struct {
	Type               SupplementalInfoItemType
	Text               *string
	RunIn              *RunIn
	VerbalIllustration *VerbalIllustration
}
