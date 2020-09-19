package types

import "github.com/pkg/errors"

// RunIn https://dictionaryapi.com/products/json#sec-2.ri
type RunIn SequenceMapping

// RunInItemType is an enum type for the types of items in RunIn
type RunInItemType int

// Values for RuninItemType
const (
	RunInItemTypeUnknown RunInItemType = iota
	RunInItemTypeText
	RunInItemTypeRunInWrap
)

// RunInItemTypeFromString returns a RunInItemType from its string ID
func RunInItemTypeFromString(id string) RunInItemType {
	switch id {
	case "text":
		return RunInItemTypeText
	case "riw":
		return RunInItemTypeRunInWrap
	default:
		return RunInItemTypeUnknown
	}
}

func (t RunInItemType) String() string {
	return []string{"", "text", "riw"}[t]
}

// Contents returns a copied slice of the contents in the RunIn
func (ri RunIn) Contents() ([]RunInItem, error) {
	items := []RunInItem{}
	for _, el := range ri {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := RunInItemTypeFromString(key)
		switch typ {
		case RunInItemTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			items = append(items, RunInItem{Type: typ, Text: &out})
		case RunInItemTypeRunInWrap:
			var out RunInWrap
			err = el.UnmarshalValue(&out)
			items = append(items, RunInItem{Type: typ, RunInWrap: &out})
		default:
			err = errors.New("unknown item type in run-in")
		}
	}
	return items, nil
}

// RunInItem is an item of the RunIn container
type RunInItem struct {
	Type      RunInItemType
	Text      *string
	RunInWrap *RunInWrap
}

// RunInWrap https://dictionaryapi.com/products/json#sec-2.ri
type RunInWrap struct {
	Word string
	WithPronounciations
	WithVariants
}
