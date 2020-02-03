package gedcom_test

import (
	"github.com/elliotchance/gedcom/tag"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestSimpleNode_ChildNodes(t *testing.T) {
	node := gedcom.NewNode(tag.TagText, "", "")

	assert.Len(t, node.Nodes(), 0)
}

func TestIsNil(t *testing.T) {
	IsNil := tf.Function(t, gedcom.IsNil)

	IsNil((*gedcom.SimpleNode)(nil)).Returns(true)
	IsNil(gedcom.NewBirthNode("")).Returns(false)
	IsNil((*gedcom.NameNode)(nil)).Returns(true)
	IsNil(gedcom.NewNameNode("")).Returns(false)

	// Untyped nil is a special case that cannot be tested above.
	assert.True(t, gedcom.IsNil(nil))
}

func TestSimpleNode_Equals(t *testing.T) {
	Equals := tf.NamedFunction(t, "SimpleNode_Equals", (*gedcom.SimpleNode).Equals)

	left := []*gedcom.SimpleNode{
		(*gedcom.SimpleNode)(nil),
		gedcom.NewNode(tag.TagVersion, "", "").(*gedcom.SimpleNode),
		gedcom.NewNode(tag.TagVersion, "a", "").(*gedcom.SimpleNode),
		gedcom.NewNode(tag.TagVersion, "", "b").(*gedcom.SimpleNode),
		gedcom.NewNode(tag.TagVersion, "a", "b").(*gedcom.SimpleNode),
	}

	right := gedcom.Nodes{
		(*gedcom.SimpleNode)(nil),
		gedcom.NewNode(tag.TagVersion, "", "").(*gedcom.SimpleNode),
		gedcom.NewNode(tag.TagVersion, "a", "").(*gedcom.SimpleNode),
		gedcom.NewNode(tag.TagVersion, "", "b").(*gedcom.SimpleNode),
		gedcom.NewNode(tag.TagVersion, "a", "b").(*gedcom.SimpleNode),
		gedcom.NewNameNode(""),
	}

	const N = false
	const Y = true

	expected := [][]bool{
		{N, N, N, N, N, N},
		{N, Y, N, N, N, N},
		{N, N, Y, N, N, N},
		{N, N, N, Y, N, N},
		{N, N, N, N, Y, N},
	}

	for i, l := range left {
		for j, r := range right {
			Equals(l, r).Returns(expected[i][j])
		}
	}
}

func TestSimpleNode_Tag(t *testing.T) {
	Tag := tf.Function(t, (*gedcom.SimpleNode).Tag)

	Tag((*gedcom.SimpleNode)(nil)).Returns(tag.Tag{})
}

func TestSimpleNode_Value(t *testing.T) {
	Value := tf.Function(t, (*gedcom.SimpleNode).Value)

	Value((*gedcom.SimpleNode)(nil)).Returns("")
}

func TestSimpleNode_Pointer(t *testing.T) {
	Pointer := tf.Function(t, (*gedcom.SimpleNode).Pointer)

	Pointer((*gedcom.SimpleNode)(nil)).Returns("")
}

func TestSimpleNode_Nodes(t *testing.T) {
	Nodes := tf.Function(t, (*gedcom.SimpleNode).Nodes)

	Nodes((*gedcom.SimpleNode)(nil)).Returns((gedcom.Nodes)(nil))
}

func TestSimpleNode_String(t *testing.T) {
	String := tf.Function(t, (*gedcom.SimpleNode).String)

	String((*gedcom.SimpleNode)(nil)).Returns("")
}

func TestSimpleNode_GEDCOMString(t *testing.T) {
	root := gedcom.NewDocument().AddIndividual("P1",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("6 MAY 1989"),
		),
	)

	assert.Equal(t, root.GEDCOMString(0), `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`)
}

func TestSimpleNode_GEDCOMLine(t *testing.T) {
	GEDCOMLine := tf.NamedFunction(t, "SimpleNode_GEDCOMLine",
		(*gedcom.SimpleNode).GEDCOMLine)

	GEDCOMLine(gedcom.NewNode(tag.TagBirth, "foo", "72").(*gedcom.BirthNode).SimpleNode, 0).
		Returns("0 @72@ BIRT foo")

	GEDCOMLine(gedcom.NewNode(tag.TagDeath, "bar", "baz").(*gedcom.DeathNode).SimpleNode, 3).Returns("3 @baz@ DEAT bar")

	GEDCOMLine(gedcom.NewDateNode("3 SEP 1945").SimpleNode, 2).
		Returns("2 DATE 3 SEP 1945")

	GEDCOMLine(gedcom.NewNode(tag.TagBirth, "foo", "72").(*gedcom.BirthNode).SimpleNode, -1).
		Returns("@72@ BIRT foo")
}

func TestSimpleNode_SetNodes(t *testing.T) {
	birth := gedcom.NewBirthNode("foo")
	assert.Nil(t, birth.Nodes())

	birth.SetNodes(gedcom.Nodes{
		gedcom.NewDateNode("3 SEP 1945"),
	})
	assert.Equal(t,
		gedcom.Nodes{gedcom.NewDateNode("3 SEP 1945")}, birth.Nodes())

	birth.SetNodes(nil)
	assert.Nil(t, birth.Nodes())
}
