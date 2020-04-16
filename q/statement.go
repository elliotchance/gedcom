package q

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
}

// Evaluate executes all of the expressions and returns the final result.
func (v *Statement) Evaluate(engine *Engine, mutInput interface{}) (interface{}, error) {
	var err error

	for _, expression := range v.Expressions {
		mutInput, err = expression.Evaluate(engine, mutInput, nil)

		if err != nil {
			return nil, err
		}
	}

	return mutInput, nil
}
