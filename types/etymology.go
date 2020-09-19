package types

import "github.com/pkg/errors"

// WithEtymology is a compositing type for parsing the `et` property
type WithEtymology struct {
	Etymology *Etymology `json:"et,omitempty"`
}

// Etymology https://dictionaryapi.com/products/json#sec-2.et
type Etymology SequenceMapping

// EtymologyItemType is an enum type for the types of items in Etymology
type EtymologyItemType int

// Values for EtymologyElementType
const (
	EtymologyItemTypeUnknown = iota
	EtymologyItemTypeText
	EtymologyItemTypeSupplementalInfo
)

// EtymologyItemTypeFromString returns a EtymologyItemType from its string ID
func EtymologyItemTypeFromString(id string) EtymologyItemType {
	switch id {
	case "text":
		return EtymologyItemTypeText
	case "et_snote":
		return EtymologyItemTypeSupplementalInfo
	default:
		return EtymologyItemTypeUnknown
	}
}

func (t EtymologyItemType) String() string {
	return []string{"", "text", "et_snote"}[t]
}

// Contents returns a copied slice of the contents in the Etymology
func (ety Etymology) Contents() ([]EtymologyItem, error) {
	items := []EtymologyItem{}
	for _, el := range ety {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := EtymologyItemTypeFromString(key)
		switch typ {
		case EtymologyItemTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			items = append(items, EtymologyItem{Type: typ, Text: &out})
		case EtymologyItemTypeSupplementalInfo:
			var out SupplementalInfo
			err = el.UnmarshalValue(&out)
			items = append(items, EtymologyItem{Type: typ, SupplementalInfo: &out})
		default:
			err = errors.New("unknown item type in etymology")
		}
	}
	return items, nil
}

// EtymologyItem is an item of the SI container
type EtymologyItem struct {
	Type             EtymologyItemType
	Text             *string
	SupplementalInfo *SupplementalInfo
}
