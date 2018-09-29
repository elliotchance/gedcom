package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewBaptismNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewBaptismNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.Equal(t, gedcom.TagBaptism, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}

func TestBaptismNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.BaptismNode).Dates)

	Dates((*gedcom.BaptismNode)(nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBaptismNode(nil, "", "", nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{})).
		Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})

	Dates(gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})
}

func TestBaptismNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.BaptismNode).Equals)

	n1 := gedcom.NewBaptismNode(nil, "foo", "", nil)
	n2 := gedcom.NewBaptismNode(nil, "bar", "", nil)

	// nils
	Equals((*gedcom.BaptismNode)(nil), (*gedcom.BaptismNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.BaptismNode)(nil)).Returns(false)
	Equals((*gedcom.BaptismNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode(nil, "foo", "", nil)).Returns(false)

	// All other cases are success.
	Equals(n1, n1).Returns(true)
	Equals(n1, n2).Returns(true)
}
