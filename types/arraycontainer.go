package types

import (
	"github.com/pkg/errors"
)

// ErrInvalidArrayMultiMapContainer is an error used when the array container is improperly structured
var ErrInvalidArrayMultiMapContainer = errors.New("invalid array container")

// ArrayMultiMapContainer is an abstract type for de/serializing a JSON structure shaped like [ [string, any] ]
type ArrayMultiMapContainer []MapItem

// Filter returns a slice of the items of the container that have the specified key
func (a ArrayMultiMapContainer) Filter(key string) []*MapItem {
	result := []*MapItem{}
	for _, el := range a {
		if k, _ := el.Key(); k == key {
			result = append(result, &el)
		}
	}
	return result
}

// MapItem is a key/value pair represented in a 2-item slice
type MapItem []interface{}

// Key returns the key string of it
func (it MapItem) Key() (string, error) {
	if len(it) != 2 {
		return "", errors.Wrap(ErrInvalidArrayMultiMapContainer, "underlying slice must have length 2")
	}
	if key, ok := it[0].(string); ok {
		return key, nil
	}
	return "", errors.Wrap(ErrInvalidArrayMultiMapContainer, "item[0] is not a string")
}

// Value returns the interface{} value of it
func (it MapItem) Value() (interface{}, error) {
	if len(it) != 2 {
		return "", errors.Wrap(ErrInvalidArrayMultiMapContainer, "underlying slice must have length 2")
	}
	return it[1], nil
}

// UnmarshalValue marshals then unmarshals the value of the item; used to assert/assign types to an otherwise opaque interface{} value
func (it MapItem) UnmarshalValue(output interface{}) error {
	value, err := it.Value()
	if err != nil {
		return err
	}

	return reUnmarshal(value, output)
}
