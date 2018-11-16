package q

import (
	"reflect"
)

// LengthExpr is a function. See Evaluate.
type LengthExpr struct{}

// Evaluate returns an integer with the number of items in the slice. This value
// will be 0 or more. If the input is not a slice then 1 will always be
// returned.
func (e *LengthExpr) Evaluate(engine *Engine, input interface{}) (interface{}, error) {
	in := reflect.ValueOf(input)

	if in.Kind() == reflect.Slice {
		return in.Len(), nil
	}

	return 1, nil
}
