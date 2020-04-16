package q

import (
	"reflect"
)

// CombineExpr will combine multiple slices of the same type into a single
// slice.
//
// If the slices are not the same type an error will be returned with a nil
// value.
type CombineExpr struct{}

func (e *CombineExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	if len(args) == 0 {
		return nil, nil
	}

	// Build a new mutSlice with all elements.
	firstArg, err := args[0].Evaluate(engine, input)
	if err != nil {
		return nil, err
	}

	mutSlice := reflect.MakeSlice(reflect.TypeOf(firstArg), 0, 0)
	for _, arg := range args {
		argValue, err := arg.Evaluate(engine, input)
		if err != nil {
			return nil, err
		}

		mutSlice = reflect.AppendSlice(mutSlice, reflect.ValueOf(argValue))
	}

	return mutSlice.Interface(), nil
}
