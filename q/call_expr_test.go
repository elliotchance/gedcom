package q_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

func TestCallExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "CallExpr_Evaluate", (*q.CallExpr).Evaluate)
	engine := &q.Engine{}

	args := []*q.Statement{{Expressions: []q.Expression{&q.ConstantExpr{Value: "123"}}}}

	Evaluate(&q.CallExpr{&q.LengthExpr{}, args}, engine, []int{1, 2, 3}, nil).
		Returns(3, nil)
}
