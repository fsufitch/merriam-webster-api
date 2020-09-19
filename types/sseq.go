package types

import (
	"github.com/pkg/errors"
)

// ErrInvalidSenseSequence is an error when the sense sequence is structured wrong
var ErrInvalidSenseSequence = errors.New("invalid sense sequence")

// WithSenseSequence is a compositing type for parsing the `sseq` property
type WithSenseSequence struct {
	SenseSequence *SenseSequence `json:"sseq,omitempty"`
}

// SenseSequence https://dictionaryapi.com/products/json#sec-2.vis
type SenseSequence SequenceMapping

// SenseSequenceItemType is an enum type for the types of items in SenseSequence
type SenseSequenceItemType int

// Values for SenseSequenceItemType
const (
	SenseSequenceItemTypeUnknown = iota
	SenseSequenceItemTypeSense
	SenseSequenceItemTypeAbbreviatedSense
	SenseSequenceItemTypeBindingSubstitute
	SenseSequenceItemTypeSubSequence
	SenseSequenceItemTypeParenthesizedSequence
)

// SenseSequenceItemTypeFromString returns a SenseSequenceItemType from its string ID
func SenseSequenceItemTypeFromString(id string) SenseSequenceItemType {
	switch id {
	case "sense":
		return SenseSequenceItemTypeSense
	case "sen":
		return SenseSequenceItemTypeAbbreviatedSense
	case "bs":
		return SenseSequenceItemTypeBindingSubstitute
	case "pseq":
		return SenseSequenceItemTypeParenthesizedSequence
	default:
		return SenseSequenceItemTypeUnknown
	}
}

func (t SenseSequenceItemType) String() string {
	return []string{"", "sense", "sen", "bs", "", "pseq"}[t]
}

// Contents returns a copied slice of the contents in the SenseSequence
func (vi SenseSequence) Contents() ([]SenseSequenceItem, error) {

	items := []SenseSequenceItem{}
	for _, el := range vi {
		if len(el) == 0 {
			return nil, errors.New("zero length item")
		}
		if _, ok := el[0].([]interface{}); ok {
			// the elements might be either proper [string, interface{}] sequence items, or [sequence, ...] sub-sequences
			// this block handles the latter case
			out := SenseSequence{}
			if err := reUnmarshal(el, &out); err != nil {
				return nil, errors.Wrap(err, "error composing subsequence")

			}
			items = append(items, SenseSequenceItem{Type: SenseSequenceItemTypeSubSequence, SubSequence: &out})
			continue
		}

		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := SenseSequenceItemTypeFromString(key)
		switch typ {
		case SenseSequenceItemTypeSense:
			var out Sense
			err = el.UnmarshalValue(&out)
			items = append(items, SenseSequenceItem{Type: typ, Sense: &out})
		case SenseSequenceItemTypeAbbreviatedSense:
			var out AbbreviatedSense
			err = el.UnmarshalValue(&out)
			items = append(items, SenseSequenceItem{Type: typ, AbbreviatedSense: &out})
		case SenseSequenceItemTypeBindingSubstitute:
			var out BindingSubstitute
			err = el.UnmarshalValue(&out)
			items = append(items, SenseSequenceItem{Type: typ, BindingSubstitute: &out})
		case SenseSequenceItemTypeParenthesizedSequence:
			var out SenseSequence
			err = el.UnmarshalValue(&out)
			items = append(items, SenseSequenceItem{Type: typ, ParenthesizedSequence: &out})
		default:
			err = errors.New("unknown element type in verbal illustration")
		}
	}
	return items, nil
}

// SenseSequenceItem is an item of the SI container
type SenseSequenceItem struct {
	Type                  SenseSequenceItemType
	Sense                 *Sense
	AbbreviatedSense      *AbbreviatedSense
	BindingSubstitute     *BindingSubstitute
	SubSequence           *SenseSequence
	ParenthesizedSequence *SenseSequence
}
