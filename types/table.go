package types

// WithTable is a compositing type for parsing the `table` property
type WithTable struct {
	// https://dictionaryapi.com/products/json#sec-2.table
	Table *Table `json:"table,omitempty"`
}

// Table https://dictionaryapi.com/products/json#sec-2.table
type Table struct {
	ID          string `json:"tableid,omitempty"`
	DisplayName string `json:"displayname,omitempty"`
}
