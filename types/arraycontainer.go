package types

import (
	"github.com/pkg/errors"
)

// ErrInvalidSequenceMapping is an error used when the sequence mapping is improperly structured
var ErrInvalidSequenceMapping = errors.New("invalid sequence mapping")

// SequenceMapping is an abstract type for de/serializing a JSON structure shaped like [ [string, any] ]
type SequenceMapping []SequenceMappingItem

// Filter returns a slice of the items of the sequence mapping that have the specified key
func (sm SequenceMapping) Filter(key string) []*SequenceMappingItem {
	result := []*SequenceMappingItem{}
	for _, it := range sm {
		if k, _ := it.Key(); k == key {
			result = append(result, &it)
		}
	}
	return result
}

// SequenceMappingItem is a key/value pair represented in a 2-item slice
type SequenceMappingItem []interface{}

// Key returns the key string of it
func (it SequenceMappingItem) Key() (string, error) {
	if len(it) != 2 {
		return "", errors.Wrap(ErrInvalidSequenceMapping, "underlying slice must have length 2")
	}
	if key, ok := it[0].(string); ok {
		return key, nil
	}
	return "", errors.Wrap(ErrInvalidSequenceMapping, "item[0] is not a string")
}

// Value returns the interface{} value of it
func (it SequenceMappingItem) Value() (interface{}, error) {
	if len(it) != 2 {
		return "", errors.Wrap(ErrInvalidSequenceMapping, "underlying slice must have length 2")
	}
	return it[1], nil
}

// UnmarshalValue marshals then unmarshals the value of the item; used to assert/assign types to an otherwise opaque interface{} value
func (it SequenceMappingItem) UnmarshalValue(output interface{}) error {
	value, err := it.Value()
	if err != nil {
		return err
	}

	return reUnmarshal(value, output)
}
