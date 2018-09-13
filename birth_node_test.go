package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewBirthNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewBirthNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.Equal(t, gedcom.TagBirth, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}

func TestBirthNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.BirthNode).Dates)

	Dates((*gedcom.BirthNode)(nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBirthNode(nil, "", "", nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBirthNode(nil, "", "", []gedcom.Node{})).
		Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})

	Dates(gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})
}

func TestBirthNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.BirthNode).Equals)

	n1 := gedcom.NewBirthNode(nil, "foo", "", nil)
	n2 := gedcom.NewBirthNode(nil, "bar", "", nil)

	// nils
	Equals((*gedcom.BirthNode)(nil), (*gedcom.BirthNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.BirthNode)(nil)).Returns(false)
	Equals((*gedcom.BirthNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode(nil, "foo", "", nil)).Returns(false)

	// All other cases are success.
	Equals(n1, n1).Returns(true)
	Equals(n1, n2).Returns(true)
}
