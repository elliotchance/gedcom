package q_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
	"sort"
)

func TestEngine_Start(t *testing.T) {
	Start := tf.Function(t, (*q.Engine).Evaluate)

	parser := q.NewParser()

	document := gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		}),
		gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
			gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
		}),
	})

	documents := []*gedcom.Document{document}

	engine, err := parser.ParseString("")
	assert.Nil(t, engine)
	assert.EqualError(t, err, "expected expression")

	engine, err = parser.ParseString(".Individuals")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns(gedcom.IndividualNodes{
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
			}),
			gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
				gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
			}),
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
			gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name | .String")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]string{
			"Elliot Chance",
			"Dina Wyche",
		}, nil)
	}

	engine, err = parser.ParseString("Names are .Individuals | .Name; Names | .String")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]string{
			"Elliot Chance",
			"Dina Wyche",
		}, nil)
	}

	engine, err = parser.ParseString("Names is .Individuals | .Name; Names")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
			gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
		}, nil)
	}

	engine, err = parser.ParseString("Names is .Individuals | .Name; BadVariable")
	if assert.NoError(t, err) {
		Start(engine, documents).Errors("no such variable BadVariable")
	}

	engine, err = parser.ParseString("?")
	if assert.NoError(t, err) {
		choices := append(documentChoices, functionAndVariableChoices...)
		choices = append(choices, "Document1")
		sort.Strings(choices)
		Start(engine, documents).Returns(choices, nil)
	}

	engine, err = parser.ParseString("Names is .Individuals | .Name; ?")
	if assert.NoError(t, err) {
		choices := append(documentChoices, functionAndVariableChoices...)
		choices = append(choices, "Document1", "Names")
		sort.Strings(choices)
		Start(engine, documents).Returns(choices, nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name; ?")
	if assert.NoError(t, err) {
		choices := append(documentChoices, functionAndVariableChoices...)
		choices = append(choices, "Document1")
		sort.Strings(choices)
		Start(engine, documents).Returns(choices, nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name | First(1)")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name | Last(23)")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
			gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | {}")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]map[string]interface{}{
			{},
			{},
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | { name: .Name }")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]map[string]interface{}{
			{"name": gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil)},
			{"name": gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil)},
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | { name: .Name, age: .Age }")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]map[string]interface{}{
			{
				"name": gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
				"age":  gedcom.Age{},
			},
			{
				"name": gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
				"age":  gedcom.Age{},
			},
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | { name: .Name | .String, age: .Age }")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]map[string]interface{}{
			{
				"name": "Elliot Chance",
				"age":  gedcom.Age{},
			},
			{
				"name": "Dina Wyche",
				"age":  gedcom.Age{},
			},
		}, nil)
	}

	engine, err = parser.ParseString("Combine(.Individuals, .Individuals) | { name: .Name | .String }")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]map[string]interface{}{
			{"name": "Elliot Chance"},
			{"name": "Dina Wyche"},
			{"name": "Elliot Chance"},
			{"name": "Dina Wyche"},
		}, nil)
	}

	engine, err = parser.ParseString("Combine(Document1 | .Individuals, Document2 | .Individuals) | { name: .Name | .String }")
	if assert.NoError(t, err) {
		Start(engine, []*gedcom.Document{document, document}).Returns([]map[string]interface{}{
			{"name": "Elliot Chance"},
			{"name": "Dina Wyche"},
			{"name": "Elliot Chance"},
			{"name": "Dina Wyche"},
		}, nil)
	}
}
