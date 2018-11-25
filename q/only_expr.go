package q

import (
	"errors"
	"reflect"
)

// OnlyExpr is a function. See Evaluate.
type OnlyExpr struct{}

// Evaluate returns a new slice that only contains the entities that have
// returned true from the condition.
func (e *OnlyExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	in := reflect.ValueOf(input)

	if len(args) != 1 {
		return nil, errors.New("function Only() must take a single argument")
	}

	if in.Kind() != reflect.Slice {
		return nil, nil
	}

	results := reflect.MakeSlice(reflect.SliceOf(TypeOfSliceElement(input)), 0, 0)

	condition := args[0]
	for i := 0; i < in.Len(); i++ {
		result, err := condition.Evaluate(engine, in.Index(i).Interface())
		if err != nil {
			return nil, err
		}

		if shouldAppend, ok := result.(bool); ok && shouldAppend {
			results = reflect.Append(results, in.Index(i))
		}
	}

	return results.Interface(), nil
}
