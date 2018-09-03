package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewEventNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewEventNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.Equal(t, gedcom.TagEvent, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}

func TestEventNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.EventNode).Dates)

	Dates(gedcom.NewEventNode(nil, "", "", nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewEventNode(nil, "", "", []gedcom.Node{})).
		Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewEventNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})

	Dates(gedcom.NewEventNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode(nil, "7 Jan 2001", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 2001", "", nil),
	})
}

func TestEventNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.EventNode).Equals)

	n1 := gedcom.NewEventNode(nil, "foo", "", nil)
	n2 := gedcom.NewEventNode(nil, "bar", "", nil)
	n3 := gedcom.NewEventNode(nil, "bar", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
		gedcom.NewDateNode(nil, "Oct 1943", "", nil),
	})
	n4 := gedcom.NewEventNode(nil, "bar", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "Oct 1943", "", nil),
	})

	// nils
	Equals((*gedcom.EventNode)(nil), (*gedcom.EventNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.EventNode)(nil)).Returns(false)
	Equals((*gedcom.EventNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode(nil, "foo", "", nil)).Returns(false)

	// General cases.
	Equals(n1, n1).Returns(false)
	Equals(n1, n2).Returns(false)
	Equals(n1, n3).Returns(false)
	Equals(n3, n4).Returns(true)
}

func TestEventNode_Years(t *testing.T) {
	Years := tf.Function(t, (*gedcom.EventNode).Years)

	Years(gedcom.NewEventNode(nil, "", "", nil)).Returns(0.0)

	Years(gedcom.NewEventNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 SEP 1943", "", nil),
	})).Returns(1943.672131147541)

	Years(gedcom.NewEventNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 SEP 1943", "", nil),
		gedcom.NewDateNode(nil, "3 SEP 1920", "", nil),
	})).Returns(1920.6730245231608)
}
