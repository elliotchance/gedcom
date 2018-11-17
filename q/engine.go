package q

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

// Engine is the compiled query. It is able to evaluate the entire query.
type Engine struct {
	Statements []*Statement
}

// Evaluate executes all of the expressions and returns the final result.
func (e *Engine) Evaluate(document *gedcom.Document) (interface{}, error) {
	for _, statement := range e.Statements {
		_, err := statement.Evaluate(e, document)

		if err != nil {
			return nil, err
		}
	}

	lastStatement := e.Statements[len(e.Statements)-1]

	return lastStatement.Evaluate(e, document)
}

func (e *Engine) StatementByVariableName(name string) (*Statement, error) {
	for _, variable := range e.Statements {
		if variable.VariableName == name {
			return variable, nil
		}
	}

	return nil, fmt.Errorf("no such variable %s", name)
}
