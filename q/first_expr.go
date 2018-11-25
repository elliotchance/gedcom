package q

import (
	"fmt"
	"reflect"
	"strconv"
)

// FirstExpr is a function. See Evaluate.
type FirstExpr struct{}

// Evaluate returns up to the number of elements in a slice.
//
// If the input value is not a slice then it is converted into a slice of one
// element before evaluating. This means that the result will always be a slice.
// The only exception to this is if the input is nil, then the result will also
// be nil.
//
// There must be exactly one argument and it must be 0 or greater. If the number
// is greater than the length of the slice all elements are returned.
func (e *FirstExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	in := reflect.ValueOf(input)

	if len(args) != 1 {
		return nil, fmt.Errorf("function First() must take a single argument")
	}

	if input == nil {
		return nil, nil
	}

	// Convert into a slice if needed.
	if in.Kind() != reflect.Slice {
		s := reflect.MakeSlice(reflect.SliceOf(in.Type()), 1, 1)
		s.Index(0).Set(in)
		in = reflect.ValueOf(s.Interface())
	}

	if in.IsNil() {
		return nil, nil
	}

	result, err := args[0].Evaluate(engine, input)
	if err != nil {
		return nil, err
	}

	max, err := strconv.Atoi(fmt.Sprintf("%v", result))
	if err != nil {
		return nil, err
	}

	if len := in.Len(); max >= len {
		max = len
	}

	return in.Slice(0, max).Interface(), nil
}
