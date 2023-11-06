package q_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

func TestConstantExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "ConstantExpr_Evaluate", (*q.ConstantExpr).Evaluate)
	engine := &q.Engine{}

	Evaluate(&q.ConstantExpr{Value: "0"}, engine, nil, nil).Returns("0", nil)
	Evaluate(&q.ConstantExpr{Value: "12.3"}, engine, nil, nil).Returns("12.3", nil)
	Evaluate(&q.ConstantExpr{Value: "foo bar"}, engine, nil, nil).Returns("foo bar", nil)
}
