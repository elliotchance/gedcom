package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewBurialNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewBurialNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.Equal(t, gedcom.TagBurial, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}

func TestBurialNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.BurialNode).Dates)

	Dates((*gedcom.BurialNode)(nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBurialNode(nil, "", "", nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBurialNode(nil, "", "", []gedcom.Node{})).
		Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBurialNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})

	Dates(gedcom.NewBurialNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})
}

func TestBurialNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.BurialNode).Equals)

	n1 := gedcom.NewBurialNode(nil, "foo", "", nil)
	n2 := gedcom.NewBurialNode(nil, "bar", "", nil)

	// nils
	Equals((*gedcom.BurialNode)(nil), (*gedcom.BurialNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.BurialNode)(nil)).Returns(false)
	Equals((*gedcom.BurialNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode(nil, "foo", "", nil)).Returns(false)

	// All other cases are success.
	Equals(n1, n1).Returns(true)
	Equals(n1, n2).Returns(true)
}
