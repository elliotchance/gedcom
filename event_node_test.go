package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewEventNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewEventNode("foo", child)

	assert.Equal(t, gedcom.TagEvent, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}

func TestEventNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.EventNode).Dates)

	Dates((*gedcom.EventNode)(nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewEventNode("")).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewEventNode("",
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("3 Sep 2001"),
	})

	Dates(gedcom.NewEventNode("",
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	})
}

func TestEventNode_Equals(t *testing.T) {
	Equals := tf.NamedFunction(t, "EventNode_Equals", (*gedcom.EventNode).Equals)

	n1 := gedcom.NewEventNode("foo")
	n2 := gedcom.NewEventNode("bar")
	n3 := gedcom.NewEventNode("bar",
		gedcom.NewDateNode("3 Sep 1943"),
		gedcom.NewDateNode("Oct 1943"),
	)
	n4 := gedcom.NewEventNode("bar",
		gedcom.NewDateNode("Oct 1943"),
	)
	n5 := gedcom.NewEventNode("",
		gedcom.NewNode(gedcom.TagType, "Domicilie", ""),
		gedcom.NewPlaceNode("Washington, District of Columbia, District of Columbia, United States"),
	)

	// nils
	Equals((*gedcom.EventNode)(nil), (*gedcom.EventNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.EventNode)(nil)).Returns(false)
	Equals((*gedcom.EventNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode("foo")).Returns(false)

	// General cases.
	Equals(n1, n1).Returns(true)
	Equals(n1, n2).Returns(false)
	Equals(n1, n3).Returns(false)
	Equals(n3, n4).Returns(true)
	Equals(n5, n5).Returns(true)
}

func TestEventNode_Years(t *testing.T) {
	Years := tf.Function(t, (*gedcom.EventNode).Years)

	Years((*gedcom.EventNode)(nil)).Returns(0.0)

	Years(gedcom.NewEventNode("")).Returns(0.0)

	Years(gedcom.NewEventNode("",
		gedcom.NewDateNode("3 SEP 1943"),
	)).Returns(1943.672131147541)

	Years(gedcom.NewEventNode("",
		gedcom.NewDateNode("3 SEP 1943"),
		gedcom.NewDateNode("3 SEP 1920"),
	)).Returns(1920.6730245231608)
}
