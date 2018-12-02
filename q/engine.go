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
//
// Evaluate expects that there is at least one document provided.
func (e *Engine) Evaluate(documents []*gedcom.Document) (interface{}, error) {
	// Before we begin we will setup the Document variables. Each document, in
	// order will be given Document1, Document2, ...
	for i, document := range documents {
		// Always prepend. The order of these variables does not matter, but we
		// need to make sure they do not land on the end where they will attempt
		// to be evaluated as the result.
		e.Statements = append([]*Statement{{
			VariableName: fmt.Sprintf("Document%d", i+1),
			Expressions: []Expression{
				&ValueExpr{Value: document},
			},
		}}, e.Statements...)
	}

	firstDocument := documents[0]

	for _, statement := range e.Statements {
		_, err := statement.Evaluate(e, firstDocument)

		if err != nil {
			return nil, err
		}
	}

	lastStatement := e.Statements[len(e.Statements)-1]

	return lastStatement.Evaluate(e, firstDocument)
}

func (e *Engine) StatementByVariableName(name string) (*Statement, error) {
	for _, variable := range e.Statements {
		if variable.VariableName == name {
			return variable, nil
		}
	}

	return nil, fmt.Errorf("no such variable %s", name)
}
