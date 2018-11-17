package q

import (
	"errors"
	"reflect"
)

// LastExpr is a function. See Evaluate.
type LastExpr struct{}

// Evaluate returns up to the number of last elements in a slice.
//
// If the input value is not a slice then it is converted into a slice of one
// element before evaluating. This means that the result will always be a slice.
// The only exception to this is if the input is nil, then the result will also
// be nil.
//
// There must be exactly one argument and it must be 0 or greater. If the number
// is greater than the length of the slice all elements are returned.
func (e *LastExpr) Evaluate(engine *Engine, input interface{}, args []interface{}) (interface{}, error) {
	in := reflect.ValueOf(input)

	if len(args) != 1 {
		return nil, errors.New("function Last() must take a single argument")
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

	len := in.Len()

	if args[0].(int) == 0 {
		return in.Slice(0, 0).Interface(), nil
	}

	start := len - args[0].(int)
	end := len - args[0].(int) + 1

	if start < 0 {
		start = 0
		end = len
	}

	return in.Slice(start, end).Interface(), nil
}
