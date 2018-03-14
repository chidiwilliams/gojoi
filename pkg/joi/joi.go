package joi

import (
	"errors"
)

// Error definitions
var (
	ErrSchemaNil = errors.New("Schema cannot be nil")
)

// Any ...
func Any() *AnySchema {
	return NewAnySchema()
}

// String ...
func String() *StringSchema {
	return NewStringSchema()
}

// Array ...
func Array() *ArraySchema {
	return NewArraySchema()
}

// Validate ...
func Validate(value interface{}, schema Schema) error {
	if !IsSet(schema) {
		return ErrSchemaNil
	}
	return schema.Root().Validate(value)
}
