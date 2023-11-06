package q_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

func TestMergeDocumentsAndIndividualsExpr_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "MergeDocumentsAndIndividualsExpr_Evaluate", (*q.MergeDocumentsAndIndividualsExpr).Evaluate)
	engine := &q.Engine{}

	argNil := []*q.Statement{}
	arg1 := []*q.Statement{{Expressions: []q.Expression{
		&q.ValueExpr{Value: gedcom.NewDocument()},
	}}}
	arg2 := []*q.Statement{
		{Expressions: []q.Expression{
			&q.ValueExpr{Value: gedcom.NewDocument()},
		}},
		{Expressions: []q.Expression{
			&q.ValueExpr{Value: gedcom.NewDocument()},
		}},
	}

	Evaluate(&q.MergeDocumentsAndIndividualsExpr{}, engine, nil, nil).
		Returns(nil, errors.New("MergeDocumentsAndIndividuals must take two arguments"))

	Evaluate(&q.MergeDocumentsAndIndividualsExpr{}, engine, nil, argNil).
		Returns(nil, errors.New("MergeDocumentsAndIndividuals must take two arguments"))

	Evaluate(&q.MergeDocumentsAndIndividualsExpr{}, engine, nil, arg1).
		Returns(nil, errors.New("MergeDocumentsAndIndividuals must take two arguments"))

	Evaluate(&q.MergeDocumentsAndIndividualsExpr{}, engine, nil, arg2).
		Returns(gedcom.NewDocument(), nil)
}
