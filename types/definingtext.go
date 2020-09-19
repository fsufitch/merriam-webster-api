package types

import "github.com/pkg/errors"

// WithDefiningText is a compositing type for parsing the `dt` property
type WithDefiningText struct {
	// https://dictionaryapi.com/products/json#sec-2.dt
	DefiningText DefiningText `json:"dt"`
}

// DefiningTextItemType is an enum type for the types of items in DefiningText
type DefiningTextItemType int

// Values for DefiningTextItemType
const (
	DefiningTextItemTypeUnknown DefiningTextItemType = iota
	DefiningTextItemTypeText
	DefiningTextItemTypeBiography
	DefiningTextItemTypeCalledAlso
	DefiningTextItemTypeRunIn
	DefiningTextItemTypeSupplementalInfo
	DefiningTextItemTypeUsageNotes
	DefiningTextItemTypeVerbalIllustrations
)

// DefiningTextItemTypeFromString returns a DefiningTextItemType from its string ID
func DefiningTextItemTypeFromString(id string) DefiningTextItemType {
	switch id {
	case "text":
		return DefiningTextItemTypeText
	case "bnw":
		return DefiningTextItemTypeBiography
	case "ca":
		return DefiningTextItemTypeCalledAlso
	case "ri":
		return DefiningTextItemTypeRunIn
	case "snote":
		return DefiningTextItemTypeSupplementalInfo
	case "uns":
		return DefiningTextItemTypeUsageNotes
	case "vis":
		return DefiningTextItemTypeVerbalIllustrations
	default:
		return DefiningTextItemTypeUnknown
	}
}

func (t DefiningTextItemType) String() string {
	return []string{"", "text", "bnw", "ca", "ri", "snote", "uns", "vis"}[t]
}

// DefiningText https://dictionaryapi.com/products/json#sec-2.dt
type DefiningText SequenceMapping

// Contents returns a copied slice of the contents in the DefiningText
func (dt DefiningText) Contents() ([]DefiningTextItem, error) {
	items := []DefiningTextItem{}
	for _, el := range dt {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := DefiningTextItemTypeFromString(key)
		switch typ {
		case DefiningTextItemTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			items = append(items, DefiningTextItem{Type: typ, Text: &out})
		case DefiningTextItemTypeBiography:
			var out Biography
			err = el.UnmarshalValue(&out)
			items = append(items, DefiningTextItem{Type: typ, Biography: &out})
		case DefiningTextItemTypeCalledAlso:
			var out CalledAlso
			err = el.UnmarshalValue(&out)
			items = append(items, DefiningTextItem{Type: typ, CalledAlso: &out})
		case DefiningTextItemTypeRunIn:
			var out RunIn
			err = el.UnmarshalValue(&out)
			items = append(items, DefiningTextItem{Type: typ, RunIn: &out})
		case DefiningTextItemTypeSupplementalInfo:
			var out SupplementalInfo
			err = el.UnmarshalValue(&out)
			items = append(items, DefiningTextItem{Type: typ, SupplementalInfo: &out})
		case DefiningTextItemTypeUsageNotes:
			var out []UsageNote
			err = el.UnmarshalValue(&out)
			items = append(items, DefiningTextItem{Type: typ, withUsageNotes: withUsageNotes{out}})
		case DefiningTextItemTypeVerbalIllustrations:
			var out []VerbalIllustration
			err = el.UnmarshalValue(&out)
			items = append(items, DefiningTextItem{Type: typ, WithVerbalIllustrations: WithVerbalIllustrations{out}})
		default:
			err = errors.New("unknown item type in defining text")
		}
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}

// DefiningTextItem is an item in the DefiningText container.
// Type indicated which property is populated with data.
type DefiningTextItem struct {
	Type             DefiningTextItemType
	Text             *string
	Biography        *Biography
	CalledAlso       *CalledAlso
	RunIn            *RunIn
	SupplementalInfo *SupplementalInfo
	withUsageNotes
	WithVerbalIllustrations
}

// Biography https://dictionaryapi.com/products/json#sec-2.bnw
type Biography struct {
	PersonalName  string `json:"pname"`
	Surname       string `json:"sname"`
	AlternateName string `json:"altname"`
}

// CalledAlso https://dictionaryapi.com/products/json#sec-2.ca
type CalledAlso struct {
	Intro   string             `json:"intro"`
	Targets []CalledAlsoTarget `json:"cats"`
}

// CalledAlsoTarget https://dictionaryapi.com/products/json#sec-2.ca
type CalledAlsoTarget struct {
	Text                string `json:"cat"`
	TargetID            string `json:"catref"`
	ParenthesizedNumber string `json:"pn"`
	WithPronounciations
	WithParenthesizedSubjectStatusLabel
}
