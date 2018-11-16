package q_test

import (
	"github.com/elliotchance/tf"
	"testing"
	"github.com/elliotchance/gedcom/q"
)

func TestLengthExpr_Evaluate(t *testing.T) {
	Evaluate := tf.Function(t, (*q.LengthExpr).Evaluate)
	engine := &q.Engine{}

	Evaluate(&q.LengthExpr{}, engine, nil).Returns(1, nil)
	Evaluate(&q.LengthExpr{}, engine, []int{1, 2, 3}).Returns(3, nil)
	Evaluate(&q.LengthExpr{}, engine, "foo bar").Returns(1, nil)
}
