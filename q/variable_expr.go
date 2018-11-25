package q

type VariableExpr struct {
	Name string
}

func (e *VariableExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	v, err := engine.StatementByVariableName(e.Name)
	if err != nil {
		return nil, err
	}

	return v.Evaluate(engine, input)
}
