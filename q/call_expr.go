package q

// CallExpr calls a function.
type CallExpr struct {
	Function Expression
	Args     []*Statement
}

func (e *CallExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	return e.Function.Evaluate(engine, input, e.Args)
}
