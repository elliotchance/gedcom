package main

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
)

func TestEngine_Start(t *testing.T) {
	Start := tf.Function(t, (*Engine).Evaluate)

	parser := NewParser()
	emptyEngine := parser.ParseString("")
	individualsEngine := parser.ParseString(".Individuals")
	individualNameEngine := parser.ParseString(".Individuals | .Name")
	individualNameStringEngine := parser.ParseString(".Individuals | .Name | .String")

	document := gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		}),
		gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
			gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
		}),
	})

	Start(emptyEngine, document).Returns(document, nil)

	Start(individualsEngine, document).Returns(gedcom.IndividualNodes{
		gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		}),
		gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
			gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
		}),
	}, nil)

	Start(individualNameEngine, document).Returns([]*gedcom.NameNode{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
	}, nil)

	Start(individualNameStringEngine, document).Returns([]string{
		"Elliot Chance",
		"Dina Wyche",
	}, nil)
}
