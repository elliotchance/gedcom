package q_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

func TestBinaryExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "BinaryExpr_Evaluate", (*q.BinaryExpr).Evaluate)
	engine := &q.Engine{}

	for _, test := range []struct {
		left     q.Expression
		op       string
		right    q.Expression
		expected bool
	}{
		{&q.ConstantExpr{"0"}, "=", &q.ConstantExpr{""}, false},
		{&q.ConstantExpr{"1.50"}, "=", &q.ConstantExpr{"1.5"}, true},
		{&q.ConstantExpr{"1.6"}, "=", &q.ConstantExpr{"1.601"}, false},
		{&q.ConstantExpr{"foo"}, "=", &q.ConstantExpr{"foo"}, true},
		{&q.ConstantExpr{"foo"}, "=", &q.ConstantExpr{"Foo"}, true},
		{&q.ConstantExpr{"\nfoo "}, "=", &q.ConstantExpr{" Foo\t"}, true},
		{&q.ConstantExpr{"foo"}, "=", &q.ConstantExpr{"bar"}, false},

		{&q.ConstantExpr{"0"}, "!=", &q.ConstantExpr{""}, true},
		{&q.ConstantExpr{"1.50"}, "!=", &q.ConstantExpr{"1.5"}, false},
		{&q.ConstantExpr{"1.6"}, "!=", &q.ConstantExpr{"1.601"}, true},
		{&q.ConstantExpr{"foo"}, "!=", &q.ConstantExpr{"foo"}, false},
		{&q.ConstantExpr{"foo"}, "!=", &q.ConstantExpr{"Foo"}, false},
		{&q.ConstantExpr{"\nfoo "}, "!=", &q.ConstantExpr{" Foo\t"}, false},
		{&q.ConstantExpr{"foo"}, "!=", &q.ConstantExpr{"bar"}, true},

		{&q.ConstantExpr{"0"}, ">", &q.ConstantExpr{""}, true},
		{&q.ConstantExpr{"1.50"}, ">", &q.ConstantExpr{"1.5"}, false},
		{&q.ConstantExpr{"1.6"}, ">", &q.ConstantExpr{"1.601"}, false},
		{&q.ConstantExpr{"foo"}, ">", &q.ConstantExpr{"foo"}, false},
		{&q.ConstantExpr{"foo"}, ">", &q.ConstantExpr{"Foo"}, false},
		{&q.ConstantExpr{"\nfoo "}, ">", &q.ConstantExpr{" Foo\t"}, false},
		{&q.ConstantExpr{"foo"}, ">", &q.ConstantExpr{"bar"}, true},

		{&q.ConstantExpr{"0"}, ">=", &q.ConstantExpr{""}, true},
		{&q.ConstantExpr{"1.50"}, ">=", &q.ConstantExpr{"1.5"}, true},
		{&q.ConstantExpr{"1.6"}, ">=", &q.ConstantExpr{"1.601"}, false},
		{&q.ConstantExpr{"foo"}, ">=", &q.ConstantExpr{"foo"}, true},
		{&q.ConstantExpr{"foo"}, ">=", &q.ConstantExpr{"Foo"}, true},
		{&q.ConstantExpr{"\nfoo "}, ">=", &q.ConstantExpr{" Foo\t"}, true},
		{&q.ConstantExpr{"foo"}, ">=", &q.ConstantExpr{"bar"}, true},

		{&q.ConstantExpr{"0"}, "<", &q.ConstantExpr{""}, false},
		{&q.ConstantExpr{"1.50"}, "<", &q.ConstantExpr{"1.5"}, false},
		{&q.ConstantExpr{"1.6"}, "<", &q.ConstantExpr{"1.601"}, true},
		{&q.ConstantExpr{"foo"}, "<", &q.ConstantExpr{"foo"}, false},
		{&q.ConstantExpr{"foo"}, "<", &q.ConstantExpr{"Foo"}, false},
		{&q.ConstantExpr{"\nfoo "}, "<", &q.ConstantExpr{" Foo\t"}, false},
		{&q.ConstantExpr{"foo"}, "<", &q.ConstantExpr{"bar"}, false},

		{&q.ConstantExpr{"0"}, "<=", &q.ConstantExpr{""}, false},
		{&q.ConstantExpr{"1.50"}, "<=", &q.ConstantExpr{"1.5"}, true},
		{&q.ConstantExpr{"1.6"}, "<=", &q.ConstantExpr{"1.601"}, true},
		{&q.ConstantExpr{"foo"}, "<=", &q.ConstantExpr{"foo"}, true},
		{&q.ConstantExpr{"foo"}, "<=", &q.ConstantExpr{"Foo"}, true},
		{&q.ConstantExpr{"\nfoo "}, "<=", &q.ConstantExpr{" Foo\t"}, true},
		{&q.ConstantExpr{"foo"}, "<=", &q.ConstantExpr{"bar"}, false},

		// TODO: Some ideas?
		//
		// >> ends with
		// !>>
		// << starts with
		// !<<
		// >< contains
		// !><
		// ~ match regexp
		// !~
	} {
		Evaluate(&q.BinaryExpr{
			Left:     test.left,
			Operator: test.op,
			Right:    test.right,
		}, engine, nil, nil).Returns(test.expected, nil)
	}
}
