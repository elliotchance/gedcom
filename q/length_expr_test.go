package q_test

import (
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/tf"
	"testing"
)

func TestLengthExpr_Evaluate(t *testing.T) {
	Evaluate := tf.Function(t, (*q.LengthExpr).Evaluate)
	engine := &q.Engine{}

	Evaluate(&q.LengthExpr{}, engine, nil, nil).Returns(1, nil)
	Evaluate(&q.LengthExpr{}, engine, []int{1, 2, 3}, nil).Returns(3, nil)
	Evaluate(&q.LengthExpr{}, engine, "foo bar", nil).Returns(1, nil)
}
