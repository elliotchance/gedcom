package q_test

import (
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/tf"
	"testing"
)

type MyStruct struct {
	Property int
}

func (ms *MyStruct) Foo() string {
	return "bar"
}

func (ms MyStruct) Baz() []string {
	return []string{"qux", "quux"}
}

func TestAccessorExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "Accessor_Evaluate", (*q.AccessorExpr).Evaluate)
	engine := &q.Engine{}

	ms1 := &MyStruct{Property: 123}
	ms2 := MyStruct{Property: 456}

	Evaluate(&q.AccessorExpr{Query: ".Foo"}, engine, ms1, nil).
		Returns("bar", nil)
	Evaluate(&q.AccessorExpr{Query: ".Foo"}, engine, ms2, nil).
		Returns("bar", nil)

	Evaluate(&q.AccessorExpr{Query: ".Baz"}, engine, ms1, nil).
		Returns([]string{"qux", "quux"}, nil)
	Evaluate(&q.AccessorExpr{Query: ".Baz"}, engine, ms2, nil).
		Returns([]string{"qux", "quux"}, nil)

	Evaluate(&q.AccessorExpr{Query: ".Property"}, engine, ms1, nil).
		Returns(123, nil)
	Evaluate(&q.AccessorExpr{Query: ".Property"}, engine, ms2, nil).
		Returns(456, nil)

	Evaluate(&q.AccessorExpr{Query: ".Missing"}, engine, ms1, nil).
		Errors(`MyStruct does not have a method or property named "Missing"`)
	Evaluate(&q.AccessorExpr{Query: ".Missing"}, engine, ms2, nil).
		Errors(`MyStruct does not have a method or property named "Missing"`)
}
