package main

import (
	"github.com/elliotchance/gedcom"
	"fmt"
)

// Engine is the compiled query. It is able to evaluate the entire query.
type Engine struct {
	Variables []*Variable
}

func NewEngine() *Engine {
	return &Engine{
		Variables: []*Variable{},
	}
}

// Evaluate executes all of the expressions and returns the final result.
func (e *Engine) Evaluate(document *gedcom.Document) (interface{}, error) {
	for _, variable := range e.Variables {
		_, err := variable.Evaluate(e, document)

		if err != nil {
			return nil, err
		}
	}

	lastVariable := e.Variables[len(e.Variables)-1]

	return lastVariable.Result, nil
}

func (e *Engine) VariableByName(name string) (*Variable, error) {
	for _, variable := range e.Variables {
		if variable.Name == name {
			return variable, nil
		}
	}

	return nil, fmt.Errorf("no such variable %s", name)
}
