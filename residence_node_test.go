package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewResidenceNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewResidenceNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.Equal(t, gedcom.TagResidence, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}

func TestResidenceNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.ResidenceNode).Dates)

	Dates(gedcom.NewResidenceNode(nil, "", "", nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{})).
		Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})

	Dates(gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})
}

func TestResidenceNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.ResidenceNode).Equals)

	n1 := gedcom.NewResidenceNode(nil, "foo", "", nil)
	n2 := gedcom.NewResidenceNode(nil, "bar", "", nil)
	n3 := gedcom.NewResidenceNode(nil, "bar", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
		gedcom.NewDateNode(nil, "Oct 1943", "", nil),
	})
	n4 := gedcom.NewResidenceNode(nil, "bar", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "Oct 1943", "", nil),
	})

	// nils
	Equals((*gedcom.ResidenceNode)(nil), (*gedcom.ResidenceNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.ResidenceNode)(nil)).Returns(false)
	Equals((*gedcom.ResidenceNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode(nil, "foo", "", nil)).Returns(false)

	// General cases.
	Equals(n1, n1).Returns(false)
	Equals(n1, n2).Returns(false)
	Equals(n1, n3).Returns(false)
	Equals(n3, n4).Returns(true)
}
