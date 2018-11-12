package main

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestEngine_Start(t *testing.T) {
	Start := tf.Function(t, (*Engine).Evaluate)

	parser := NewParser()

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
			"Length",
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
			"Length",
			"Names",
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
			"Length",
		}, nil)
	}
}
