package q_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

func TestFirstExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "FirstExpr_Evaluate", (*q.FirstExpr).Evaluate)
	engine := &q.Engine{}

	err := errors.New("function First() must take a single argument")

	argNil := []*q.Statement{}
	arg0 := []*q.Statement{{Expressions: []q.Expression{&q.ConstantExpr{Value: "0"}}}}
	arg1 := []*q.Statement{{Expressions: []q.Expression{&q.ConstantExpr{Value: "1"}}}}
	arg2 := []*q.Statement{{Expressions: []q.Expression{&q.ConstantExpr{Value: "2"}}}}
	arg5 := []*q.Statement{{Expressions: []q.Expression{&q.ConstantExpr{Value: "5"}}}}

	Evaluate(&q.FirstExpr{}, engine, nil, arg0).Returns(nil, nil)
	Evaluate(&q.FirstExpr{}, engine, nil, arg1).Returns(nil, nil)
	Evaluate(&q.FirstExpr{}, engine, nil, arg5).Returns(nil, nil)
	Evaluate(&q.FirstExpr{}, engine, nil, argNil).Returns(nil, err)

	Evaluate(&q.FirstExpr{}, engine, ([]int)(nil), arg0).Returns(nil, nil)
	Evaluate(&q.FirstExpr{}, engine, ([]string)(nil), arg1).Returns(nil, nil)
	Evaluate(&q.FirstExpr{}, engine, ([]MyStruct)(nil), arg5).Returns(nil, nil)
	Evaluate(&q.FirstExpr{}, engine, ([]string)(nil), argNil).Returns(nil, err)

	Evaluate(&q.FirstExpr{}, engine, []int{1, 2, 3}, arg0).
		Returns([]int{}, nil)
	Evaluate(&q.FirstExpr{}, engine, []int{1, 2, 3}, arg1).
		Returns([]int{1}, nil)
	Evaluate(&q.FirstExpr{}, engine, []int{1, 2, 3}, arg2).
		Returns([]int{1, 2}, nil)
	Evaluate(&q.FirstExpr{}, engine, []int{1, 2, 3}, arg5).
		Returns([]int{1, 2, 3}, nil)
	Evaluate(&q.FirstExpr{}, engine, []int{1, 2, 3}, argNil).
		Returns(nil, err)

	Evaluate(&q.LengthExpr{}, engine, "foo bar", arg0).
		Returns([]string{}, nil)
	Evaluate(&q.LengthExpr{}, engine, "foo bar", arg1).
		Returns([]string{"foo bar"}, nil)
	Evaluate(&q.LengthExpr{}, engine, "foo bar", arg5).
		Returns([]string{"foo bar"}, nil)
	Evaluate(&q.LengthExpr{}, engine, "foo bar", argNil).
		Returns(nil, err)
}
