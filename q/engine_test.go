package q_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
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

	engine, err := parser.ParseString("")
	assert.Nil(t, engine)
	assert.EqualError(t, err, "expected expression")

	engine, err = parser.ParseString(".Individuals")
	if assert.NoError(t, err) {
		Start(engine, document).Returns(gedcom.IndividualNodes{
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
		Start(engine, document).Returns([]*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
			gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name | .String")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]string{
			"Elliot Chance",
			"Dina Wyche",
		}, nil)
	}

	engine, err = parser.ParseString("Names are .Individuals | .Name; Names | .String")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]string{
			"Elliot Chance",
			"Dina Wyche",
		}, nil)
	}

	engine, err = parser.ParseString("Names is .Individuals | .Name; Names")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
			gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
		}, nil)
	}

	engine, err = parser.ParseString("Names is .Individuals | .Name; BadVariable")
	if assert.NoError(t, err) {
		Start(engine, document).Errors("no such variable BadVariable")
	}

	engine, err = parser.ParseString("?")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]string{
			".AddNode",
			".Families",
			".Individuals",
			".NodeByPointer",
			".Nodes",
			".Places",
			".Sources",
			".String",
			"?",
			"First",
			"Last",
			"Length",
			"Only",
		}, nil)
	}

	engine, err = parser.ParseString("Names is .Individuals | .Name; ?")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]string{
			".AddNode",
			".Families",
			".Individuals",
			".NodeByPointer",
			".Nodes",
			".Places",
			".Sources",
			".String",
			"?",
			"First",
			"Last",
			"Length",
			"Names",
			"Only",
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name; ?")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]string{
			".AddNode",
			".Families",
			".Individuals",
			".NodeByPointer",
			".Nodes",
			".Places",
			".Sources",
			".String",
			"?",
			"First",
			"Last",
			"Length",
			"Only",
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name | First(1)")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | .Name | Last(23)")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
			gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | {}")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]map[string]interface{}{
			{},
			{},
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | { name: .Name }")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]map[string]interface{}{
			{"name": gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil)},
			{"name": gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil)},
		}, nil)
	}

	engine, err = parser.ParseString(".Individuals | { name: .Name, age: .Age }")
	if assert.NoError(t, err) {
		Start(engine, document).Returns([]map[string]interface{}{
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
		Start(engine, document).Returns([]map[string]interface{}{
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
}
