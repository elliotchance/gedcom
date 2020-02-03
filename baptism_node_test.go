package gedcom_test

import (
	"github.com/elliotchance/gedcom/tag"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewBaptismNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewBaptismNode("foo", child)

	assert.Equal(t, tag.TagBaptism, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}

func TestBaptismNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.BaptismNode).Dates)

	Dates((*gedcom.BaptismNode)(nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBaptismNode("")).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBaptismNode("")).
		Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBaptismNode("",
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("3 Sep 2001"),
	})

	Dates(gedcom.NewBaptismNode("",
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	})
}

func TestBaptismNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.BaptismNode).Equals)

	n1 := gedcom.NewBaptismNode("foo")
	n2 := gedcom.NewBaptismNode("bar")

	// nils
	Equals((*gedcom.BaptismNode)(nil), (*gedcom.BaptismNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.BaptismNode)(nil)).Returns(false)
	Equals((*gedcom.BaptismNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode("foo")).Returns(false)

	// All other cases are success.
	Equals(n1, n1).Returns(true)
	Equals(n1, n2).Returns(true)
}
