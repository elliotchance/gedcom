package q_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

func TestCombineExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "CombineExpr_Evaluate",
		(*q.CombineExpr).Evaluate)
	engine := &q.Engine{}

	slice1 := []int{3, 4, 5}
	slice2 := []int{1, 2}

	Evaluate(&q.CombineExpr{}, engine, nil, []*q.Statement{}).Returns(nil, nil)

	Evaluate(&q.CombineExpr{}, engine, nil, []*q.Statement{
		{Expressions: []q.Expression{&q.ValueExpr{slice2}}},
	}).Returns([]int{1, 2}, nil)

	Evaluate(&q.CombineExpr{}, engine, nil, []*q.Statement{
		{Expressions: []q.Expression{&q.ValueExpr{slice1}}},
		{Expressions: []q.Expression{&q.ValueExpr{slice2}}},
	}).Returns([]int{3, 4, 5, 1, 2}, nil)
}
