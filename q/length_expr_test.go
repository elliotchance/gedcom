package q_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

func TestLengthExpr_Evaluate(t *testing.T) {
	Evaluate := tf.Function(t, (*q.LengthExpr).Evaluate)
	engine := &q.Engine{}

	Evaluate(&q.LengthExpr{}, engine, nil, nil).Returns(1, nil)
	Evaluate(&q.LengthExpr{}, engine, []int{1, 2, 3}, nil).Returns(3, nil)
	Evaluate(&q.LengthExpr{}, engine, "foo bar", nil).Returns(1, nil)
}
