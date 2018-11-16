package q

import "github.com/elliotchance/gedcom"

// Statement represents a single discreet operation in the engine.
type Statement struct {
	// VariableName must be unique amongst other variables and must not be the
	// name of an existing function. The name is also allow to be empty which
	// means that the result cannot be referenced in other expressions.
	VariableName string

	// Expressions are separated by pipes. The result of each evaluated
	// expressions is used as the input to the next expressions. The input value
	// for the first expression is the gedcom.Document.
	Expressions []Expression

	// isEvaluated prevents Evaluate from calculating the result again when the
	// variable is retrieved. Once Evaluate is complete with a success or
	// failure the results is store in the "result" and "error" properties.
	isEvaluated bool

	// result and error hold the cached result from Evaluate. You should not
	// access these directly, but instead it's safe to called Evaluate many
	// times.
	result interface{}
	error error
}

// Evaluate executes all of the expressions and returns the final result.
func (v *Statement) Evaluate(engine *Engine, document *gedcom.Document) (interface{}, error) {
	if v.isEvaluated {
		return v.result, v.error
	}

	defer func() {
		v.isEvaluated = true
	}()

	v.result = document

	for _, expression := range v.Expressions {
		v.result, v.error = expression.Evaluate(engine, v.result)

		if v.error != nil {
			return nil, v.error
		}
	}

	return v.result, nil
}
