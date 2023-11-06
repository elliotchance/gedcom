package q_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

func TestObjectExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "ObjectExpr_Evaluate", (*q.ObjectExpr).Evaluate)

	engine := &q.Engine{}
	ms1 := &MyStruct{Property: 123}

	Evaluate(&q.ObjectExpr{}, engine, nil, nil).
		Returns(map[string]interface{}{}, nil)

	Evaluate(&q.ObjectExpr{
		Data: map[string]*q.Statement{
			"foo": {
				Expressions: []q.Expression{
					&q.AccessorExpr{Query: ".Foo"},
				},
			},
		},
	}, engine, ms1, nil).Returns(map[string]interface{}{
		"foo": "bar",
	}, nil)

	Evaluate(&q.ObjectExpr{
		Data: map[string]*q.Statement{
			"foo": {
				Expressions: []q.Expression{
					&q.AccessorExpr{Query: ".Foo"},
				},
			},
			"bar": {
				Expressions: []q.Expression{
					&q.AccessorExpr{Query: ".Foo"},
				},
			},
		},
	}, engine, nil, nil).Returns(map[string]interface{}{
		"foo": nil,
		"bar": nil,
	}, nil)
}
