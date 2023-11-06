package q_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

var documentChoices = []string{
	".AddFamily",
	".AddFamilyWithHusbandAndWife",
	".AddIndividual",
	".AddNode",
	".DeleteNode",
	".Families",
	".GEDCOMString",
	".Individuals",
	".NodeByPointer",
	".Nodes",
	".Places",
	".SetNodes",
	".Sources",
	".String",
	".Warnings",
}

var functionAndVariableChoices = []string{
	"?",
	"Combine",
	"First",
	"Last",
	"Length",
	"MergeDocumentsAndIndividuals",
	"NodesWithTagPath",
	"Only",
}

func TestQuestionMarkExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "QuestionMarkExpr_Evaluate",
		(*q.QuestionMarkExpr).Evaluate)
	engine := &q.Engine{}

	expected := append([]string{".Baz", ".Foo", ".Panic"},
		functionAndVariableChoices...)

	Evaluate(&q.QuestionMarkExpr{}, engine, &MyStruct{}, nil).
		Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, MyStruct{}, nil).
		Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, []*MyStruct{{}, {}}, nil).
		Returns(expected, nil)
	Evaluate(&q.QuestionMarkExpr{}, engine, []*MyStruct{}, nil).
		Returns(expected, nil)
}
