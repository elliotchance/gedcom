package gedcom_test

import (
	"github.com/elliotchance/gedcom/tag"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewBurialNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewBurialNode("foo", child)

	assert.Equal(t, tag.TagBurial, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}

func TestBurialNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.BurialNode).Dates)

	Dates((*gedcom.BurialNode)(nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBurialNode("")).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBurialNode("",
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("3 Sep 2001"),
	})

	Dates(gedcom.NewBurialNode("",
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	})
}

func TestBurialNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.BurialNode).Equals)

	n1 := gedcom.NewBurialNode("foo")
	n2 := gedcom.NewBurialNode("bar")

	// nils
	Equals((*gedcom.BurialNode)(nil), (*gedcom.BurialNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.BurialNode)(nil)).Returns(false)
	Equals((*gedcom.BurialNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode("foo")).Returns(false)

	// All other cases are success.
	Equals(n1, n1).Returns(true)
	Equals(n1, n2).Returns(true)
}
