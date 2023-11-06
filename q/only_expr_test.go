package q_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

func TestOnlyExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "OnlyExpr_Evaluate", (*q.OnlyExpr).Evaluate)
	engine := &q.Engine{}

	err := errors.New("function Only() must take a single argument")

	argNil := []*q.Statement{}
	argEq := []*q.Statement{{Expressions: []q.Expression{&q.BinaryExpr{
		Left:     &q.AccessorExpr{".Property"},
		Operator: "=",
		Right:    &q.ConstantExpr{"52"},
	}}}}
	argBadComp := []*q.Statement{{Expressions: []q.Expression{
		&q.ConstantExpr{Value: "foo"},
	}}}

	Evaluate(&q.OnlyExpr{}, engine, nil, argEq).Returns(nil, nil)
	Evaluate(&q.OnlyExpr{}, engine, nil, argBadComp).Returns(nil, nil)
	Evaluate(&q.OnlyExpr{}, engine, nil, argNil).Returns(nil, err)

	Evaluate(&q.OnlyExpr{}, engine, MyStruct{Property: 52}, argEq).Returns(nil, nil)
	Evaluate(&q.OnlyExpr{}, engine, MyStruct{Property: 52}, argBadComp).Returns(nil, nil)
	Evaluate(&q.OnlyExpr{}, engine, MyStruct{Property: 13}, argEq).Returns(nil, nil)
	Evaluate(&q.OnlyExpr{}, engine, MyStruct{}, argNil).Returns(nil, err)

	Evaluate(&q.OnlyExpr{}, engine, []MyStruct{{Property: 52}}, argEq).Returns([]MyStruct{{Property: 52}}, nil)
	Evaluate(&q.OnlyExpr{}, engine, []MyStruct{{Property: 52}}, argBadComp).Returns([]MyStruct{}, nil)
	Evaluate(&q.OnlyExpr{}, engine, []MyStruct{{Property: 13}}, argEq).Returns([]MyStruct{}, nil)
	Evaluate(&q.OnlyExpr{}, engine, []MyStruct{{}}, argNil).Returns(nil, err)
}
