package q_test

import (
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/tf"
	"testing"
)

var documentChoices = []string{
	".AddNode",
	".Families",
	".GEDCOMString",
	".Individuals",
	".NodeByPointer",
	".Nodes",
	".Places",
	".SetNodes",
	".Sources",
	".String",
}

var functionAndVariableChoices = []string{
	"?",
	"Combine",
	"First",
	"Last",
	"Length",
	"MergeDocumentsAndIndividuals",
	"Only",
}

func TestQuestionMarkExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "QuestionMarkExpr_Evaluate",
		(*q.QuestionMarkExpr).Evaluate)
	engine := &q.Engine{}

	expected := append([]string{".Baz", ".Foo"}, functionAndVariableChoices...)

	Evaluate(&q.QuestionMarkExpr{}, engine, &MyStruct{}, nil).
		Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, MyStruct{}, nil).
		Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, []*MyStruct{{}, {}}, nil).
		Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, []*MyStruct{}, nil).
		Returns(expected, nil)
}
