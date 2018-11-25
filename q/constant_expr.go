package q

// ConstantExpr represents a floating-point number or string.
type ConstantExpr struct {
	Value string
}

func (e *ConstantExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	return e.Value, nil
}
