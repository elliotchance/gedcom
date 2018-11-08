package main

import "github.com/elliotchance/gedcom"

// Engine is the compiled query. It is able to evaluate the entire query.
type Engine struct {
	// Expressions are separated by pipes. The result of each evaluated
	// expressions is used as the input to the next expressions. The input value
	// for the first expression is the gedcom.Document.
	Expressions []Expression
}

// Evaluate executes all of the expressions and returns the final result.
func (e *Engine) Evaluate(document *gedcom.Document) (interface{}, error) {
	var result interface{} = document
	var err error

	for _, expression := range e.Expressions {
		result, err = expression.Evaluate(result)

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
