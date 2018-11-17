package q_test

import (
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/tf"
	"testing"
)

func TestCallExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "CallExpr_Evaluate", (*q.CallExpr).Evaluate)
	engine := &q.Engine{}

	Evaluate(&q.CallExpr{&q.LengthExpr{}, []interface{}{123}}, engine, []int{1, 2, 3}, nil).
		Returns(3, nil)
}
