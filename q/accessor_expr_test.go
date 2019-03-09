package q_test

import (
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/tf"
	"testing"
	"github.com/stretchr/testify/assert"
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

func (ms MyStruct) Panic() string {
	panic("oh no!")
}

func TestAccessorExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "Accessor_Evaluate", (*q.AccessorExpr).Evaluate)
	engine := &q.Engine{}

	ms1 := &MyStruct{Property: 123}
	ms2 := MyStruct{Property: 456}

	Evaluate(&q.AccessorExpr{Query: ".Foo"}, engine, nil, nil).
		Returns(nil, nil)
	Evaluate(&q.AccessorExpr{Query: ".Foo"}, engine, (*MyStruct)(nil), nil).
		Returns("bar", nil)

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

	// Panics
	q := &q.AccessorExpr{Query: ".Panic"}
	_, err := q.Evaluate(engine, ms1, nil)
	assert.NotNil(t, err)
	assert.Regexp(t, `^panic MyStruct.Panic: oh no!\ngoroutine `, err)
}
