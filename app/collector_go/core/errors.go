package core

import (
	"fmt"
	"reflect"
)

type UnmarshalTypeError struct {
	Value string
	Type  reflect.Type
	Err   error
}

func (e *UnmarshalTypeError) Error() string {
	return fmt.Sprintf("unmarshalfailed, cannot unmarshal %s into Go value type %s",
		e.Value, e.Type.String())
}

type ErrInvalidFacility struct {
	Facility Facility
	Err      error
}

func (e *ErrInvalidFacility) Error() string {
	return fmt.Sprintf("facility id and facility index id must be the same: %v", e.Facility)
}
