package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
)

func TestDeepEqual(t *testing.T) {
	DeepEqual := tf.Function(t, gedcom.DeepEqual)

	// These two are the same.
	n1 := gedcom.NewResidenceNode("",
		gedcom.NewDateNode("3 SEP 1987"),
	)
	n2 := gedcom.NewResidenceNode("",
		gedcom.NewDateNode("3 SEP 1987"),
	)

	// Different variations.
	n3 := gedcom.NewResidenceNode("",
		gedcom.NewDateNode("5 SEP 1987"),
	)
	n4 := gedcom.NewResidenceNode("",
		gedcom.NewDateNode("5 SEP 1987"),
		gedcom.NewDateNode("3 SEP 1987"),
	)
	n5 := gedcom.NewResidenceNode("",
		gedcom.NewDateNode("3 SEP 1987"),
		gedcom.NewDateNode("5 SEP 1987"),
	)

	// More complex examples.
	n6 := gedcom.NewDocument().AddIndividual("P1",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewResidenceNode("",
			gedcom.NewDateNode("3 SEP 1987"),
			gedcom.NewPlaceNode("England"),
		),
	)
	n7 := gedcom.NewDocument().AddIndividual("P1",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewResidenceNode("",
			gedcom.NewDateNode("3 SEP 1987"),
			gedcom.NewPlaceNode("England"),
		),
	)
	n8 := gedcom.NewDocument().AddIndividual("P1",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewResidenceNode("",
			gedcom.NewDateNode("3 SEP 1987"),
			gedcom.NewPlaceNode("London, England"),
		),
	)

	// Nils.
	DeepEqual((*gedcom.SimpleNode)(nil), (*gedcom.SimpleNode)(nil)).Returns(false)
	DeepEqual((*gedcom.SimpleNode)(nil), n1).Returns(false)
	DeepEqual(n1, (*gedcom.SimpleNode)(nil)).Returns(false)

	// Equal.
	DeepEqual(n1, n1).Returns(true) // #4
	DeepEqual(n1, n2).Returns(true)
	DeepEqual(n2, n1).Returns(true)
	DeepEqual(n1, n3).Returns(false)
	DeepEqual(n1, n3).Returns(false)

	// Different amount of children.
	DeepEqual(n3, n4).Returns(false) // #9

	// Deep equal.
	DeepEqual(n4, n5).Returns(true) // #10
	DeepEqual(n5, n4).Returns(true)
	DeepEqual(n6, n7).Returns(true)
	DeepEqual(n7, n6).Returns(true)
	DeepEqual(n7, n8).Returns(false)
}

func TestDeepEqualNodes(t *testing.T) {
	DeepEqualNodes := tf.Function(t, gedcom.DeepEqualNodes)

	// These two are the same.
	n1 := gedcom.NewResidenceNode("",
		gedcom.NewDateNode("3 SEP 1987"),
	)
	n2 := gedcom.NewResidenceNode("",
		gedcom.NewDateNode("3 SEP 1987"),
	)
	n3 := gedcom.NewResidenceNode("",
		gedcom.NewDateNode("5 SEP 1987"),
	)

	// There aren't too many tests here because the fiddly stuff is handled in
	// the tests for DeepEqual.

	DeepEqualNodes(nil, nil).Returns(true)
	DeepEqualNodes(gedcom.Nodes{n1}, gedcom.Nodes{n1}).Returns(true)
	DeepEqualNodes(gedcom.Nodes{n1}, gedcom.Nodes{n2}).Returns(true)
	DeepEqualNodes(gedcom.Nodes{n1, n2}, gedcom.Nodes{n1, n2}).Returns(true)

	DeepEqualNodes(gedcom.Nodes{n1, n2}, gedcom.Nodes{n1}).Returns(false)
	DeepEqualNodes(gedcom.Nodes{n1}, gedcom.Nodes{n3}).Returns(false)
}
