package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestNewResidenceNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewResidenceNode("foo", child)

	assert.Equal(t, gedcom.TagResidence, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}

func TestResidenceNode_Dates(t *testing.T) {
	Dates := tf.Function(t, (*gedcom.ResidenceNode).Dates)

	Dates((*gedcom.ResidenceNode)(nil)).Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewResidenceNode("")).
		Returns([]*gedcom.DateNode(nil))

	Dates(gedcom.NewResidenceNode("",
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("3 Sep 2001"),
	})

	Dates(gedcom.NewResidenceNode("",
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	)).Returns([]*gedcom.DateNode{
		gedcom.NewDateNode("7 Jan 2001"),
		gedcom.NewDateNode("3 Sep 2001"),
	})
}

func TestResidenceNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.ResidenceNode).Equals)

	n1 := gedcom.NewResidenceNode("foo")
	n2 := gedcom.NewResidenceNode("bar")
	n3 := gedcom.NewResidenceNode("bar",
		gedcom.NewDateNode("3 Sep 1943"),
		gedcom.NewDateNode("Oct 1943"),
	)
	n4 := gedcom.NewResidenceNode("bar",
		gedcom.NewDateNode("Oct 1943"),
	)

	// nils
	Equals((*gedcom.ResidenceNode)(nil), (*gedcom.ResidenceNode)(nil)).Returns(false)
	Equals(n1, (*gedcom.ResidenceNode)(nil)).Returns(false)
	Equals((*gedcom.ResidenceNode)(nil), n1).Returns(false)

	// Wrong node type.
	Equals(n1, gedcom.NewNameNode("foo")).Returns(false)

	// General cases.
	Equals(n1, n1).Returns(true)
	Equals(n1, n2).Returns(true)
	Equals(n1, n3).Returns(false)
	Equals(n3, n4).Returns(true)

	// Residences are equal if they contain the same place but both do not
	// specify a date.
	r1 := gedcom.NewResidenceNode("",
		gedcom.NewSourceNode("@S619368194@", ""),
		gedcom.NewPlaceNode("Worcestershire"),
	)
	r2 := gedcom.NewResidenceNode("",
		gedcom.NewPlaceNode("Worcestershire"),
		gedcom.NewSourceNode("@S619368193@", "", gedcom.NewNode(gedcom.TagFromString("_APID"), "1,2056::540816", "")),
	)
	Equals(r1, r2).Returns(true)
}

func TestResidenceNode_Years(t *testing.T) {
	Years := tf.Function(t, (*gedcom.ResidenceNode).Years)

	Years(gedcom.NewResidenceNode("")).Returns(0.0)

	Years(gedcom.NewResidenceNode("",
		gedcom.NewDateNode("3 SEP 1943"),
	)).Returns(1943.672131147541)

	Years(gedcom.NewResidenceNode("",
		gedcom.NewDateNode("3 SEP 1943"),
		gedcom.NewDateNode("3 SEP 1920"),
	)).Returns(1920.6730245231608)
}
