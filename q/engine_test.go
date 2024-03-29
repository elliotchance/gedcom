package q_test

import (
	"testing"

	"sort"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestEngine_Start(t *testing.T) {
	Start := tf.Function(t, (*q.Engine).Evaluate)

	parser := q.NewParser()

	document := gedcom.NewDocument()
	document.AddIndividual("P1",
		gedcom.NewNameNode("Elliot /Chance/"),
	)
	document.AddIndividual("P2",
		gedcom.NewNameNode("Dina /Wyche/"),
	)

	documents := []*gedcom.Document{document}

	engine, err := parser.ParseString("")
	assert.Nil(t, engine)
	assert.EqualError(t, err, "expected expression")

	engine, err = parser.ParseString(".Individuals")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns(document.Individuals(), nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]*gedcom.NameNode{
			gedcom.NewNameNode("Elliot /Chance/"),
			gedcom.NewNameNode("Dina /Wyche/"),
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
			gedcom.NewNameNode("Elliot /Chance/"),
			gedcom.NewNameNode("Dina /Wyche/"),
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
			gedcom.NewNameNode("Elliot /Chance/"),
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name | Last(23)")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]*gedcom.NameNode{
			gedcom.NewNameNode("Elliot /Chance/"),
			gedcom.NewNameNode("Dina /Wyche/"),
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
			{"name": gedcom.NewNameNode("Elliot /Chance/")},
			{"name": gedcom.NewNameNode("Dina /Wyche/")},
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | { name: .Name, age: .Age }")
	if assert.NoError(t, err) {
		Start(engine, documents).Returns([]map[string]interface{}{
			{
				"name": gedcom.NewNameNode("Elliot /Chance/"),
				"age":  gedcom.Age{},
			},
			{
				"name": gedcom.NewNameNode("Dina /Wyche/"),
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
