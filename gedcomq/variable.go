package main

import "github.com/elliotchance/gedcom"

type Variable struct {
	// Name must be unique amongst other variable and must not be the name of an
	// existing function. The name is also allow to be empty which means that
	// the result cannot be referenced in other expressions.
	Name string

	// Expressions are separated by pipes. The result of each evaluated
	// expressions is used as the input to the next expressions. The input value
	// for the first expression is the gedcom.Document.
	Expressions []Expression

	IsEvaluated bool

	Result interface{}
	Error error
}

// Evaluate executes all of the expressions and returns the final result.
func (v *Variable) Evaluate(engine *Engine, document *gedcom.Document) (interface{}, error) {
	if v.IsEvaluated {
		return v.Result, v.Error
	}

	defer func() {
		v.IsEvaluated = true
	}()

	v.Result = document

	for _, expression := range v.Expressions {
		v.Result, v.Error = expression.Evaluate(engine, v.Result)

		if v.Error != nil {
			return nil, v.Error
		}
	}

	return v.Result, nil
}
