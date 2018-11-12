package main

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestQuestionMarkExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "QuestionMarkExpr_Evaluate", (*QuestionMarkExpr).Evaluate)
	engine := &Engine{}

	expected := []string{".Baz", ".Foo", "?", "Length"}

	Evaluate(&QuestionMarkExpr{}, engine, &MyStruct{}).Returns(expected, nil)
	Evaluate(&QuestionMarkExpr{}, engine, MyStruct{}).Returns(expected, nil)
	Evaluate(&QuestionMarkExpr{}, engine, []*MyStruct{{}, {}}).Returns(expected, nil)
	Evaluate(&QuestionMarkExpr{}, engine, []*MyStruct{}).Returns(expected, nil)
}
