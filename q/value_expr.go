package q

// ValueExpr holds a single value.
//
// It is different from ConstantExpr because it cannot be instantiated from the
// q language, but acts as a placeholder for prepared values.
type ValueExpr struct {
	Value interface{}
}

func (e *ValueExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	return e.Value, nil
}
