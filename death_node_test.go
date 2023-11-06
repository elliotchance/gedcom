package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewDeathNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewDeathNode("foo", child)

	assert.Equal(t, gedcom.TagDeath, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}

func TestDeathNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.DeathNode).Dates)

	Dates((*gedcom.DeathNode)(nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewDeathNode("")).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewDeathNode("",
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("3 Sep 2001"),
	})

	Dates(gedcom.NewDeathNode("",
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	})
}

func TestDeathNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.DeathNode).Equals)

	n1 := gedcom.NewDeathNode("foo")
	n2 := gedcom.NewDeathNode("bar")

	// nils
	Equals((*gedcom.DeathNode)(nil), (*gedcom.DeathNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.DeathNode)(nil)).Returns(false)
	Equals((*gedcom.DeathNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode("foo")).Returns(false)

	// All other cases are success.
	Equals(n1, n1).Returns(true)
	Equals(n1, n2).Returns(true)
}
