package main

import (
	"reflect"
	"sort"
)

// QuestionMark ("?") is a special function. See Evaluate.
type QuestionMark struct{}

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
func (e *QuestionMark) Evaluate(engine *Engine, input interface{}) (interface{}, error) {
	in := reflect.TypeOf(input)

	if in.Kind() == reflect.Slice {
		value := reflect.Zero(TypeOfSliceElement(input)).Interface()

		return e.Evaluate(engine, value)
	}

	if in.Kind() != reflect.Ptr {
		in = reflect.New(in).Type()
	}

	options := []string{}
	for i := 0; i < in.NumMethod(); i++ {
		methodName := "." + in.Method(i).Name
		options = append(options, methodName)
	}

	for function := range Functions {
		options = append(options, function)
	}

	sort.Strings(options)

	return options, nil
}
