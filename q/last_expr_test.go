package q_test

import (
	"errors"
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/tf"
	"testing"
)

func TestLastExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "LastExpr_Evaluate", (*q.LastExpr).Evaluate)
	engine := &q.Engine{}

	err := errors.New("function Last() must take a single argument")

	Evaluate(&q.LastExpr{}, engine, nil, []interface{}{0}).Returns(nil, nil)
	Evaluate(&q.LastExpr{}, engine, nil, []interface{}{1}).Returns(nil, nil)
	Evaluate(&q.LastExpr{}, engine, nil, []interface{}{5}).Returns(nil, nil)
	Evaluate(&q.LastExpr{}, engine, nil, []interface{}{}).Returns(nil, err)

	Evaluate(&q.LastExpr{}, engine, ([]int)(nil), []interface{}{0}).
		Returns(nil, nil)
	Evaluate(&q.LastExpr{}, engine, ([]string)(nil), []interface{}{1}).
		Returns(nil, nil)
	Evaluate(&q.LastExpr{}, engine, ([]MyStruct)(nil), []interface{}{5}).
		Returns(nil, nil)
	Evaluate(&q.LastExpr{}, engine, ([]MyStruct)(nil), []interface{}{}).
		Returns(nil, err)

	Evaluate(&q.LastExpr{}, engine, []int{1, 2, 3}, []interface{}{0}).
		Returns([]int{}, nil)
	Evaluate(&q.LastExpr{}, engine, []int{1, 2, 3}, []interface{}{1}).
		Returns([]int{3}, nil)
	Evaluate(&q.LastExpr{}, engine, []int{1, 2, 3}, []interface{}{5}).
		Returns([]int{1, 2, 3}, nil)
	Evaluate(&q.LastExpr{}, engine, []int{1, 2, 3}, []interface{}{}).
		Returns(nil, err)

	Evaluate(&q.LengthExpr{}, engine, "foo bar", []interface{}{0}).
		Returns([]string{}, nil)
	Evaluate(&q.LengthExpr{}, engine, "foo bar", []interface{}{1}).
		Returns([]string{"foo bar"}, nil)
	Evaluate(&q.LengthExpr{}, engine, "foo bar", []interface{}{5}).
		Returns([]string{"foo bar"}, nil)
	Evaluate(&q.LengthExpr{}, engine, "foo bar", []interface{}{}).
		Returns(nil, err)
}
