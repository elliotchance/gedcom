package q_test

import (
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/tf"
	"testing"
)

func TestQuestionMarkExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "QuestionMarkExpr_Evaluate", (*q.QuestionMarkExpr).Evaluate)
	engine := &q.Engine{}

	expected := []string{".Baz", ".Foo", "?", "First", "Last", "Length", "Only"}

	Evaluate(&q.QuestionMarkExpr{}, engine, &MyStruct{}, nil).
		Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, MyStruct{}, nil).
		Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, []*MyStruct{{}, {}}, nil).
		Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, []*MyStruct{}, nil).
		Returns(expected, nil)
}
