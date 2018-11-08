package main

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestQuestionMark_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "QuestionMark_Evaluate", (*QuestionMark).Evaluate)

	expected := []string{".Baz", ".Foo", "?", "Length"}

	Evaluate(&QuestionMark{}, &MyStruct{}).Returns(expected, nil)
	Evaluate(&QuestionMark{}, MyStruct{}).Returns(expected, nil)
	Evaluate(&QuestionMark{}, []*MyStruct{{}, {}}).Returns(expected, nil)
	Evaluate(&QuestionMark{}, []*MyStruct{}).Returns(expected, nil)
}
