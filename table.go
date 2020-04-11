package sqlStorage

// Based on: https://github.com/rs/rest-layer/blob/f3fd620c10308bc912e48dd643d3f1642fa30d2c/schema/string.go

import (
	"errors"
)

// String validates string based values
type Table struct {
}

// Compile compiles and validate regexp if any.
func (v *Table) Compile() (err error) {
	return
}

// ValidateQuery implements schema.FieldQueryValidator interface
func (v Table) ValidateQuery(value interface{}) (interface{}, error) {
	s, ok := value.(string)
	if !ok {
		return nil, errors.New("not a string")
	}
	return s, nil
}

// Validate validates and normalize string based value.
func (v Table) Validate(value interface{}) (interface{}, error) {
	// Pre-check that compilation was successful.
	s, ok := value.(string)
	if !ok {
		return nil, errors.New("not a string")
	}
	return s, nil
}

