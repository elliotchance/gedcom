package q_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/q"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNodesWithTagPathExpr_Evaluate(t *testing.T) {
	engine := &q.Engine{}

	doc := gedcom.NewDocument()
	individual := doc.AddIndividual("P1")
	individual.AddBirthDate("16 Apr 1973")
	individual2 := doc.AddIndividual("P2")
	individual2.AddBirthDate("8 Mar 1884")

	for testName, test := range map[string]struct {
		input          interface{}
		args           []*q.Statement
		expectedGEDCOM string
	}{
		"Nil": {
			input:          nil,
			args:           nil,
			expectedGEDCOM: "",
		},
		"IndividualNone": {
			input:          individual,
			args:           nil,
			expectedGEDCOM: "",
		},
		"IndividualBIRT": {
			input: individual,
			args: []*q.Statement{
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "BIRT"}}},
			},
			expectedGEDCOM: "0 BIRT\n1 DATE 16 Apr 1973\n",
		},
		"IndividualDATE": {
			input: individual,
			args: []*q.Statement{
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "DATE"}}},
			},
			expectedGEDCOM: "",
		},
		"IndividualBIRTDATE": {
			input: individual,
			args: []*q.Statement{
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "BIRT"}}},
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "DATE"}}},
			},
			expectedGEDCOM: "0 DATE 16 Apr 1973\n",
		},
		"IndividualBIRTDEAT": {
			input: individual,
			args: []*q.Statement{
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "BIRT"}}},
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "DEAT"}}},
			},
			expectedGEDCOM: "",
		},
		"IndividualsBIRT": {
			input: doc.Individuals(),
			args: []*q.Statement{
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "BIRT"}}},
			},
			expectedGEDCOM: "0 BIRT\n1 DATE 16 Apr 1973\n0 BIRT\n1 DATE 8 Mar 1884\n",
		},
		"IndividualsDATE": {
			input: doc.Individuals(),
			args: []*q.Statement{
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "DATE"}}},
			},
			expectedGEDCOM: "",
		},
		"IndividualsBIRTDATE": {
			input: doc.Individuals(),
			args: []*q.Statement{
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "BIRT"}}},
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "DATE"}}},
			},
			expectedGEDCOM: "0 DATE 16 Apr 1973\n0 DATE 8 Mar 1884\n",
		},
		"IndividualsBIRTDEAT": {
			input: doc.Individuals(),
			args: []*q.Statement{
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "BIRT"}}},
				{Expressions: []q.Expression{&q.ConstantExpr{Value: "DEAT"}}},
			},
			expectedGEDCOM: "",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			actual, err := (&q.NodesWithTagPathExpr{}).Evaluate(engine, test.input, test.args)
			require.NoError(t, err)

			assert.Equal(t, test.expectedGEDCOM, gedcom.NewDocumentWithNodes(actual.(gedcom.Nodes)).String())
		})
	}
}
