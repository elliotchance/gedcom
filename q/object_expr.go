package q

import "reflect"

// ObjectExpr creates an object from keys and values.
type ObjectExpr struct {
	Data map[string]*Statement
}

func (e *ObjectExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	in := reflect.ValueOf(input)

	// If it is a slice we need to Evaluate each one.
	if in.Kind() == reflect.Slice {
		mapType := reflect.TypeOf([]map[string]interface{}{})
		mutResults := reflect.MakeSlice(mapType, 0, 0)

		for i := 0; i < in.Len(); i++ {
			result, err := e.Evaluate(engine, in.Index(i).Interface(), nil)
			if err != nil {
				return nil, err
			}

			mutResults = reflect.Append(mutResults, reflect.ValueOf(result))
		}

		return mutResults.Interface(), nil
	}

	// Evaluate single.
	m := map[string]interface{}{}

	for name, stmt := range e.Data {
		value, err := stmt.Evaluate(engine, input)
		if err != nil {
			return nil, err
		}

		m[name] = value
	}

	return m, nil
}
