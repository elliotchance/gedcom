package main

type VariableExpr struct {
	VariableName string
}

func (e *VariableExpr) Evaluate(engine *Engine, input interface{}) (interface{}, error) {
	v, err := engine.VariableByName(e.VariableName)
	if err != nil {
		return nil, err
	}

	return v.Evaluate(engine, nil)
}
