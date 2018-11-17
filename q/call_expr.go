package q

// CallExpr calls a function.
type CallExpr struct {
	Function Expression
	Args     []interface{}
}

func (e *CallExpr) Evaluate(engine *Engine, input interface{}, args []interface{}) (interface{}, error) {
	return e.Function.Evaluate(engine, input, e.Args)
}
