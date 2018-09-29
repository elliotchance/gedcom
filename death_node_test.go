package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewDeathNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewDeathNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.Equal(t, gedcom.TagDeath, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}

func TestDeathNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.DeathNode).Dates)

	Dates((*gedcom.DeathNode)(nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewDeathNode(nil, "", "", nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewDeathNode(nil, "", "", []gedcom.Node{})).
		Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})

	Dates(gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})
}

func TestDeathNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.DeathNode).Equals)

	n1 := gedcom.NewDeathNode(nil, "foo", "", nil)
	n2 := gedcom.NewDeathNode(nil, "bar", "", nil)

	// nils
	Equals((*gedcom.DeathNode)(nil), (*gedcom.DeathNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.DeathNode)(nil)).Returns(false)
	Equals((*gedcom.DeathNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode(nil, "foo", "", nil)).Returns(false)

	// All other cases are success.
	Equals(n1, n1).Returns(true)
	Equals(n1, n2).Returns(true)
}
