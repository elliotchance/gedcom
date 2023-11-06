package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewBirthNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewBirthNode("foo", child)

	assert.Equal(t, gedcom.TagBirth, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}

func TestBirthNode_Dates(t *testing.T) {
	Dates := tf.NamedFunction(t, "BirthNode_Dates", (*gedcom.BirthNode).Dates)

	Dates((*gedcom.BirthNode)(nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBirthNode("")).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBirthNode("",
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("3 Sep 2001"),
	})

	Dates(gedcom.NewBirthNode("",
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	})
}

func TestBirthNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.BirthNode).Equals)

	n1 := gedcom.NewBirthNode("foo")
	n2 := gedcom.NewBirthNode("bar")

	// nils
	Equals((*gedcom.BirthNode)(nil), (*gedcom.BirthNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.BirthNode)(nil)).Returns(false)
	Equals((*gedcom.BirthNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode("foo")).Returns(false)

	// All other cases are success.
	Equals(n1, n1).Returns(true)
	Equals(n1, n2).Returns(true)
}
