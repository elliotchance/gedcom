package q_test

import (
	"github.com/elliotchance/tf"
	"testing"
	"github.com/elliotchance/gedcom/q"
)

func TestQuestionMarkExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "QuestionMarkExpr_Evaluate", (*q.QuestionMarkExpr).Evaluate)
	engine := &q.Engine{}

	expected := []string{".Baz", ".Foo", "?", "Length"}

	Evaluate(&q.QuestionMarkExpr{}, engine, &MyStruct{}).Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, MyStruct{}).Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, []*MyStruct{{}, {}}).Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, []*MyStruct{}).Returns(expected, nil)
}
