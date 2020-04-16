package q

import (
	"reflect"
	"sort"
)

// QuestionMarkExpr ("?") is a special function. See Evaluate.
type QuestionMarkExpr struct{}

// "?" is a special function that can be used to show all of the possible next
// functions and accessors. This is useful when exploring data by creating the
// query interactively.
//
// For example the following query:
//
//   .Individuals | ?
//
// Returns (most items removed for brevity):
//
//   [
//     ".AddNode",
//     ".Age",
//     ".AgeAt",
//     ...
//     ".SurroundingSimilarity",
//     ".Tag",
//     ".Value",
//     "?",
//     "Length"
//   ]
//
func (e *QuestionMarkExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	mutIn := reflect.TypeOf(input)

	if mutIn.Kind() == reflect.Slice {
		value := reflect.Zero(TypeOfSliceElement(input)).Interface()

		return e.Evaluate(engine, value, nil)
	}

	if mutIn.Kind() != reflect.Ptr {
		mutIn = reflect.New(mutIn).Type()
	}

	mutOptions := []string{}

	// Accessors
	for i := 0; i < mutIn.NumMethod(); i++ {
		methodName := "." + mutIn.Method(i).Name
		mutOptions = append(mutOptions, methodName)
	}

	// Functions
	for function := range Functions {
		mutOptions = append(mutOptions, function)
	}

	// Variables
	for _, statement := range engine.Statements {
		if statement.VariableName != "" {
			mutOptions = append(mutOptions, statement.VariableName)
		}
	}

	sort.Strings(mutOptions)

	return mutOptions, nil
}
