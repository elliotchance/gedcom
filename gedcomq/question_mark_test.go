package main

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestQuestionMark_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "QuestionMark_Evaluate", (*QuestionMark).Evaluate)
	engine := NewEngine()

	expected := []string{".Baz", ".Foo", "?", "Length"}

	Evaluate(&QuestionMark{}, engine, &MyStruct{}).Returns(expected, nil)
	Evaluate(&QuestionMark{}, engine, MyStruct{}).Returns(expected, nil)
	Evaluate(&QuestionMark{}, engine, []*MyStruct{{}, {}}).Returns(expected, nil)
	Evaluate(&QuestionMark{}, engine, []*MyStruct{}).Returns(expected, nil)
}
