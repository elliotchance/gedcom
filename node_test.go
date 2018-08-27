package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

func TestNodeGedcom(t *testing.T) {
	NodeGedcom := tf.Function(t, gedcom.NodeGedcom)

	root := gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewSimpleNode(nil, gedcom.TagBirth, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
		}),
	})

	NodeGedcom((*gedcom.SimpleNode)(nil)).Returns("")
	NodeGedcom(root).Returns(`0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`)
}
