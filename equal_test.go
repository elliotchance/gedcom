package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
)

func TestDeepEqual(t *testing.T) {
	DeepEqual := tf.Function(t, gedcom.DeepEqual)

	// These two are the same.
	n1 := gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 SEP 1987", "", nil),
	})
	n2 := gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 SEP 1987", "", nil),
	})

	// Different variations.
	n3 := gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "5 SEP 1987", "", nil),
	})
	n4 := gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "5 SEP 1987", "", nil),
		gedcom.NewDateNode(nil, "3 SEP 1987", "", nil),
	})
	n5 := gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 SEP 1987", "", nil),
		gedcom.NewDateNode(nil, "5 SEP 1987", "", nil),
	})

	// More complex examples.
	n6 := gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "3 SEP 1987", "", nil),
			gedcom.NewPlaceNode(nil, "England", "", nil),
		}),
	})
	n7 := gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "3 SEP 1987", "", nil),
			gedcom.NewPlaceNode(nil, "England", "", nil),
		}),
	})
	n8 := gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "3 SEP 1987", "", nil),
			gedcom.NewPlaceNode(nil, "London, England", "", nil),
		}),
	})

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
	n1 := gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 SEP 1987", "", nil),
	})
	n2 := gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 SEP 1987", "", nil),
	})
	n3 := gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "5 SEP 1987", "", nil),
	})

	// There aren't too many tests here because the fiddly stuff is handled in
	// the tests for DeepEqual.

	DeepEqualNodes(nil, nil).Returns(true)
	DeepEqualNodes([]gedcom.Node{n1}, []gedcom.Node{n1}).Returns(true)
	DeepEqualNodes([]gedcom.Node{n1}, []gedcom.Node{n2}).Returns(true)
	DeepEqualNodes([]gedcom.Node{n1, n2}, []gedcom.Node{n1, n2}).Returns(true)

	DeepEqualNodes([]gedcom.Node{n1, n2}, []gedcom.Node{n1}).Returns(false)
	DeepEqualNodes([]gedcom.Node{n1}, []gedcom.Node{n3}).Returns(false)
}
